
import { Component, OnInit, OnDestroy, ɵisComponentDefPendingResolution } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { InputGroupModule } from 'primeng/inputgroup';
import { InputGroupAddonModule } from 'primeng/inputgroupaddon';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';
import { CheckboxModule } from 'primeng/checkbox';
import { PanelModule } from 'primeng/panel';
import { CardModule } from 'primeng/card';
import { SidebarModule } from 'primeng/sidebar';
import { ProgressBarModule } from 'primeng/progressbar';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { FieldsetModule } from 'primeng/fieldset';
import { webSocket } from 'rxjs/webSocket'

//Services & Classes
import { QueueDownloads, DownloadMedia, VideoData, VideoDataRequest } from '../../classes/video-data';
import { DownloadService } from '../../services/download.service';
import { SimplecardComponent } from "../simplecard/simplecard.component";
import { SharedDataService, Operation } from '../../services/shared-data.service';
import { Messages, Severity, wsApiUrl } from '../../constants/messages'


interface ExtractionOptions {
    Identifier: string;
    GetAudioOnly: boolean;
    GetSubs: boolean;
}


@Component({
    selector: 'app-downloads',
    standalone: true,
    imports: [ToastModule, ProgressBarModule, FieldsetModule, ProgressSpinnerModule, SidebarModule, CardModule, FormsModule,
        InputGroupModule, InputGroupAddonModule, InputTextModule, ButtonModule, CommonModule, CheckboxModule, PanelModule,
        SimplecardComponent],
    providers: [DownloadService, MessageService, Messages],
    templateUrl: './downloads.component.html',
    styleUrl: './downloads.component.scss'
})
export class DownloadsComponent implements OnInit {

    wsApiURL: string = wsApiUrl
    wsMessage: string;
    serverLogs: string;
    queuedItems: VideoData[] = [new VideoData()]
    nilMetadata = new VideoData()

    constructor(private messageService: MessageService,
        private currentDL: DownloadService,
        private sharedData: SharedDataService,
        private msg: Messages) {
        this.wsMessage = msg.wsMessage
        this.serverLogs = msg.serverLogs
    }

    loading: boolean = false;
    sock = webSocket(this.wsApiURL)

    activeDLImage = ''
    activeDLChannel = ''
    activeDLTitle = ''


    async sendRequest() {
        this.sock.subscribe();
        this.sock.next(this.wsMessage);
        // this.sock.complete();

        this.sock.subscribe({
            next: msg => this.updateLogs(JSON.stringify(msg)),
            error: err => { this.updateLogs('{"download": "web-socket connection is closed."}') },
            // complete: () => console.log('complete')
        });

    }

    async updateLogs(message: string) {

        const log = JSON.parse(message)
        if (log.download === this.msg.downloadComplete) {
            //1. remove this item from queued items
            this.sharedData.queuedItemsMetadata = this.sharedData.queuedItemsMetadata.filter(x => x.title != this.activeDLTitle && x.channel != this.activeDLChannel)
            this.sharedData.setQueuedItemsMetadata(this.sharedData.queuedItemsMetadata, Operation.Replace)

            this.queuedItems = this.sharedData.getQueuedItemsMetadata()  //update local object

            //2. trigger download for next item in queue
            if (this.sharedData.queuedItemsMetadata.length > 0) {

                setTimeout(() => {
                    this.showMessage("New download starting in 2 seconds",
                        this.msg.Severities[Severity.danger].toLowerCase(),
                        this.msg.Severities[Severity.danger])
                    this.tgrMediaDownload();                                 //trigger the next download
                    setTimeout(() => { this.sendRequest(); }, 500);          //trigger stats checker if all goes well    
                }, 2000);

                

            } else {
                //3. Handle isDownloadActive
                this.sharedData.isDownloadActive = false;
                this.sharedData.setIsDownloadActive(false);
                this.sharedData.activeDownloadMetadata = this.sharedData.queuedItemsMetadata
                this.sharedData.setActiveDownloadMetadata(this.sharedData.activeDownloadMetadata)
            }
            //update UI logs
            this.serverLogs = log.download
        } else if (this.serverLogs === this.msg.downloadComplete) {
            this.serverLogs = this.serverLogs + ' ' + log.download
        } else if (log.download.indexOf(this.msg.downloadInfoIdentifier) !== -1) {
            this.serverLogs = log.download
        }
    }

    //css-classes
    homeBoxActive = 'home-box'
    contentBoxActive = 'content-box'

    urlPlaceholder = 'Video or Playlist URL'

    sidebarVisible: boolean = false;
    options: ExtractionOptions = { Identifier: '', GetAudioOnly: false, GetSubs: false }

    flipCheckbox(event: any, option: string): void {
        if (option === 'GetAudioOnly') {
            this.options.GetAudioOnly = !event.checked;
        } else if (option === 'GetSubs') {
            this.options.GetSubs = !event.checked;
        }
    }

    ngOnInit() {

        this.sharedData.isDownloadActive = this.sharedData.getIsDownloadActive()
        this.sharedData.queuedItemsMetadata = this.sharedData.getQueuedItemsMetadata()
        this.sharedData.activeDownloadMetadata = this.sharedData.getActiveDownloadMetadata()

        // // if there is an active download
         if (this.sharedData.isDownloadActive) {
             this.sendRequest()
             this.populateVideoMetadata()
         }

        // //if there are queued items but no active downloads
        // if (!this.sharedData.isDownloadActive && this.sharedData.queuedItemsMetadata.length > 0) {
        //     this.tgrMediaDownload();
        //     this.sendRequest();
        //     // setTimeout(() => { this.sendRequest(); }, 500);
        // }
    }

    ngOnDestroy() {
        // prevent memory leak when component destroyed
        if (this.sock.closed) {
            this.sock.unsubscribe();
        } else {
            this.sock.complete();
            this.sock.unsubscribe();
        }
    }

    async GetMedia() {
        this.loading = true;
        let metadataRequest = await this.GetMetadataRequest(this.options)
        if (metadataRequest.Indicator === '') {
            this.showMessage('No URL or Identifier provided', 'error', 'error')
            return
        }

        let metadata: VideoData[] = await this.currentDL.getMetadata(metadataRequest)

        //add item to queue
        this.sharedData.queuedItemsMetadata.push(...metadata)
        this.sharedData.setQueuedItemsMetadata(this.sharedData.queuedItemsMetadata, Operation.Replace)
        this.queuedItems = this.sharedData.queuedItemsMetadata  //update local object

        //if download in-progress add to queue
        if (this.sharedData.isDownloadActive) {            
            //Completion Process
            this.GetMediaCompleteResult()
            return
        }

        //trigger the download/add to queue
        await this.tgrMediaDownload();

        //trigger stats checker if all goes well
        setTimeout(() => { this.sendRequest(); }, 500); 

        //Completion Process
        this.GetMediaCompleteResult()
    }

    GetMediaCompleteResult() {

        //Complete Result
        this.resetDownloadOptions();
        this.loading = false;
        this.showMessage('Video/Playlist queued', 'info', 'Info')
    }

    GetMetadataRequest(options: ExtractionOptions) {

        let request = new VideoDataRequest()

        request.Indicator = options.Identifier
        request.SubtitlesReq = options.GetSubs
        request.IsAudioOnly = options.GetAudioOnly

        if (request.Indicator === '') {
            request.Indicator = 'UMBEkWFMacc'
            this.showMessage('Using Test Video Indicator', 'info', 'Info')
            return request
        }
        return request
    }

    populateVideoMetadata() {

        if (this.sharedData.activeDownloadMetadata.length == 0) {
            this.sharedData.activeDownloadMetadata = this.sharedData.getActiveDownloadMetadata()
        }

        this.activeDLChannel = this.sharedData.activeDownloadMetadata[0].channel
        this.activeDLTitle = this.sharedData.activeDownloadMetadata[0].title
        this.activeDLImage = this.sharedData.activeDownloadMetadata[0].thumbnail
    }

    async GetMediaRequest(metadata: VideoData) {

        let request = new DownloadMedia()

        request.VideoId = metadata.video_id
        request.VideoURL = metadata.original_url

        let queueDownloads: QueueDownloads = { DownloadMedia: [request] };
        return queueDownloads
    }

    async tgrMediaDownload() {

        //update to activeDL
        if (this.sharedData.activeDownloadMetadata.length == 0) {
            this.sharedData.activeDownloadMetadata.push(this.sharedData.queuedItemsMetadata[0])
            this.sharedData.setActiveDownloadMetadata(this.sharedData.activeDownloadMetadata)
        } else {
            this.sharedData.activeDownloadMetadata[0] = structuredClone(this.sharedData.queuedItemsMetadata[0])
            this.sharedData.setActiveDownloadMetadata(this.sharedData.activeDownloadMetadata)
        }

        //trigger download request
        let downloadRequest = await this.GetMediaRequest(this.sharedData.queuedItemsMetadata[0])
        if (downloadRequest.DownloadMedia[0].VideoId === -1 || downloadRequest.DownloadMedia[0].VideoURL === '') {
            this.showMessage('VideoId & VideoURL are required', 'error', 'Error')
            return
        }

        //populate downloading video UI details
        this.populateVideoMetadata()

        //remove queued item at index 0, since it is no longer queued
        this.sharedData.queuedItemsMetadata.splice(0, 1)
        this.sharedData.setQueuedItemsMetadata(this.sharedData.queuedItemsMetadata, Operation.Replace)

        //set download active
        this.sharedData.isDownloadActive = true
        this.sharedData.setIsDownloadActive(true)

        //service call
        let triggerDownload: string = await this.currentDL.getMedia(downloadRequest)

        //service call response and ui messaging
        if (triggerDownload === this.msg.triggerDownloadApiSuccessResponse) {
            //show success message
            this.showMessage(this.msg.triggerDownloadApiSuccessResponse,
                this.msg.Severities[Severity.success].toLowerCase(),
                this.msg.Severities[Severity.success])
        }
    }

    async resetDownloadOptions() {
        this.options.GetAudioOnly = false
        this.options.GetSubs = false
        this.options.Identifier = ''
    }


    //Toast Messages
    showMessage(message: string, severity: string, summary: string) {
        this.messageService.add({ severity: severity, summary: summary, detail: message });
    }
}


