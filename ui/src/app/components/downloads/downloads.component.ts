
import { Component, OnInit, OnDestroy } from '@angular/core';
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


    async getDownloadStatus() {
        this.sock.subscribe();
        this.sock.next(this.wsMessage);

        this.sock.subscribe({
            next: msg => { this.updateLogs(JSON.stringify(msg)); console.log('msg:', msg) },
            error: err => { this.updateLogs('{"download": "web-socket connection is closed."}'); console.log('err:', err) },
            complete: () => { this.wsCloseWithDownloadComplete() }
        });
        //this.sock.complete();
    }

    wsCloseWithDownloadComplete() {
        console.log('ws-close message frame recieved from server.')
        this.sharedData.isDownloadActive = false
        this.sharedData.setIsDownloadActive(false)
    }

    async updateLogs(message: string) {

        const log = JSON.parse(message)
        if (log.download === this.msg.downloadComplete) {
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
    isPlaylist: boolean = false

    flipCheckbox(event: any, option: string): void {
        if (option === 'GetAudioOnly') {
            this.options.GetAudioOnly = !event.checked;
        } else if (option === 'GetSubs') {
            this.options.GetSubs = !event.checked;
        }
    }

    async ngOnInit() {

        this.sharedData.isPlaylist = this.sharedData.getIsPlaylist()
        this.sharedData.isDownloadActive = this.sharedData.getIsDownloadActive()
        this.sharedData.activeDownloadMetadata = this.sharedData.getActiveDownloadMetadata()
        
        //get queued-items on reload
        await this.getQueuedItems()

        //if there is an active download
        if (this.sharedData.isDownloadActive) {
            this.getDownloadStatus()
            this.populateVideoMetadata()
        }
    }

    ngOnDestroy() {
        // prevent memory leak when component destroyed
        if (this.sock.closed) {
            this.sock.complete();
            this.sock.unsubscribe();
        } else {
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

        if (metadata.length > 1) {
            // this.sharedData.setIsPlaylist(true);
            metadata.forEach(item => { item.isPlaylistVideo = true })
        }

        //delta update for all videos
        let allVideos = this.sharedData.getlstVideos()
        if (allVideos !== null) {
            allVideos.push(...metadata)
        } else {
            allVideos = [];
            allVideos.push(...metadata)
        }
        this.sharedData.setlstVideos(allVideos)

        //trigger stats checker if all goes well
        setTimeout(() => { this.getDownloadStatus(); }, 500);
        //tryout await here - see what happens?
        //await this.getDownloadStatus();

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
        this.serverLogs = ">>>waiting for server logs<<<"
    }

    async resetDownloadOptions() {
        this.options.GetAudioOnly = false
        this.options.GetSubs = false
        this.options.Identifier = ''
    }

    async getQueuedItems() {
        this.sidebarVisible = true
        await this.currentDL.getQueuedItems().then(x => this.queuedItems = x)
    }


    //Toast Messages
    showMessage(message: string, severity: string, summary: string) {
        this.messageService.add({ severity: severity, summary: summary, detail: message });
    }
}


