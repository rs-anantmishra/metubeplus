
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
import { RemovePrefixPipe } from '../../utilities/remove-prefix.pipe'


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
    SimplecardComponent, RemovePrefixPipe],
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

    loading: boolean = false;
    urlInputDisabled: boolean = false;

    constructor(private messageService: MessageService,
        private svcDownload: DownloadService,
        private sharedData: SharedDataService,
        private msg: Messages) {

        this.wsMessage = msg.wsMessage
        this.serverLogs = msg.serverLogs
    }

    sock = webSocket(this.wsApiURL)

    activeDLImage = ''
    activeDLChannel = ''
    activeDLTitle = ''
    activeDownload: VideoData = new VideoData()


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
        let isActiveDownload = this.sharedData.getIsDownloadActive()

        if (!isActiveDownload) {
            await this.getAndSaveActiveDownload()
            this.populateVideoMetadata()
            if (this.activeDownload !== null && this.activeDownload.title !== '') {
                this.sharedData.setIsDownloadActive(true)
            }
        }

        const log = JSON.parse(message)
        if (log.download === this.msg.downloadComplete) {
            this.serverLogs = log.download
            this.sharedData.setIsDownloadActive(false)
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
        await this.getQueuedItems(false)

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
        this.urlInputDisabled = true;
        let metadataRequest = await this.GetMetadataRequest(this.options)
        if (metadataRequest.Indicator === '') {
            this.showMessage('No URL or Identifier provided', 'error', 'error')
            return
        }

        let metadata: VideoData[] = await this.svcDownload.getMetadata(metadataRequest)
        let isDownloadActive = this.sharedData.getIsDownloadActive()
        if (!isDownloadActive) {
            await this.getAndSaveActiveDownload()
            this.populateVideoMetadata()

            //set to true
            this.sharedData.setIsDownloadActive(true)
        }

        if (metadata.length > 1) {
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

        //Completion Process
        this.GetMediaCompleteResult()
    }

    GetMediaCompleteResult() {

        //Complete Result
        this.resetDownloadOptions();
        this.loading = false;
        this.urlInputDisabled = false;
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
        //clear previous
        this.activeDownload = new VideoData()

        //get ActiveDownloadMetadata
        this.activeDownload = this.sharedData.getActiveDownloadMetadata()[0]
        console.log('activedownload is null?', this.activeDownload)

        this.activeDLChannel = this.activeDownload.channel
        this.activeDLTitle = this.activeDownload.title
        this.activeDLImage = this.activeDownload.thumbnail
        this.serverLogs = ">>>waiting for server logs<<<"

    }

    async resetDownloadOptions() {
        this.options.GetAudioOnly = false
        this.options.GetSubs = false
        this.options.Identifier = ''
    }

    async getQueuedItems(openSidebar: boolean) {
        if (openSidebar) {
            this.sidebarVisible = true
        }
        await this.svcDownload.getQueuedItems("queued").then(item => { this.queuedItems = item })
    }

    async getAndSaveActiveDownload() {
        await this.svcDownload.getQueuedItems("downloading").then(item => { console.log(item); this.sharedData.setActiveDownloadMetadata(item); })
    }



    //Toast Messages
    showMessage(message: string, severity: string, summary: string) {
        this.messageService.add({ severity: severity, summary: summary, detail: message });
    }
}


