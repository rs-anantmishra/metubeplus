
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
import { BehaviorSubject, queue } from 'rxjs';

interface ExtractionOptions {
  Identifier: string;
  GetAudioOnly: boolean;
  GetSubs: boolean;
}


@Component({
  selector: 'app-downloads',
  standalone: true,
  imports: [ToastModule, ProgressBarModule, FieldsetModule, ProgressSpinnerModule, SidebarModule, CardModule, FormsModule, InputGroupModule, InputGroupAddonModule, InputTextModule, ButtonModule, CommonModule, CheckboxModule, PanelModule, SimplecardComponent],
  providers: [DownloadService, MessageService],
  templateUrl: './downloads.component.html',
  styleUrl: './downloads.component.scss'
})
export class DownloadsComponent implements OnInit {
  loading: boolean = false;
  sock = webSocket('ws://localhost:3000/ws/downloadstatus')
  wsMessage = 'No active downloads'

  activeDLImage = ''
  activeDLChannel = ''
  activeDLTitle = ''

  serverlogs = 'No logs available.'

  sendRequest() {
    this.sock.subscribe();
    this.sock.next(this.wsMessage);
    // this.sock.complete();
    this.sock.subscribe({
      next: msg => this.updateLogs(JSON.stringify(msg)),
      error: err => { console.log('error in ws connection', err), this.updateLogs('{"download": "web-socket connection is closed."}') },
      //complete: () => console.log('complete') 
    });
  }

  updateLogs(message: string) {
    this.serverlogs = message
    const serverlog = JSON.parse(message)
    this.serverlogs = serverlog.download

    //these below will be set from metadata only
    //
    // this.downloadingChannel = serverlog.channel
    // this.downloadingTitle = serverlog.title
  }


  constructor(private messageService: MessageService,
    private currentDL: DownloadService) { }

  interval: any;
  dlProgress: number = 98
  dl: VideoData = new VideoData()

  //css-classes
  homeBoxActive = 'home-box'
  contentBoxActive = 'content-box'

  sidebarVisible: boolean = false;
  options: ExtractionOptions = { Identifier: '', GetAudioOnly: false, GetSubs: false }

  flipCheckbox(event: any, option: string): void {
    if (option === 'GetAudioOnly') {
      this.options.GetAudioOnly = !event.checked;
    } else if (option === 'GetSubs') {
      this.options.GetSubs = !event.checked;
    }
  }

  placeholder = 'Video or Playlist URL'
  ngOnInit() {
    //this.sendRequest();
  }

  ngOnDestroy() {
  }



  async GetMedia() {

    this.loading = true;
    let metadataRequest = await this.GetMetadataRequest(this.options)
    if (metadataRequest.Indicator === '') {
      this.showMessage('No URL or Identifier provided', 'error', 'error')
      return
    }

    //this.dl = await this.currentDL.getMetadata(request)
    let metadata: VideoData[] = await this.currentDL.getMetadata(metadataRequest)
    this.populateVideoMetadata(metadata)

    //trigger download
    let downloadRequest = await this.GetMediaRequest(metadata)
    console.log('dl req', downloadRequest)
    if (downloadRequest.DownloadMedia[0].VideoId === -1 || downloadRequest.DownloadMedia[0].VideoURL === '') {
      this.showMessage('VideoId & VideoURL are required', 'error', 'Error')
      return
    }

    let triggerDownload: string = await this.currentDL.getMedia(downloadRequest)
    console.log(triggerDownload)

    //trigger stats checker if all goes well
    this.sendRequest();

    //Complete Result
    this.loading = false;
    this.showMessage('Video/Playlist queued', 'info', 'Info')
  }

  async GetMetadataRequest(options: ExtractionOptions) {

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

  populateVideoMetadata(metadata: VideoData[]) {
    this.activeDLChannel = metadata[0].channel
    this.activeDLTitle = metadata[0].title
    this.activeDLImage = metadata[0].thumbnail
  }

  async GetMediaRequest(metadata: VideoData[]) {

    let request = new DownloadMedia()
    
    request.VideoId = metadata[0].video_id
    request.VideoURL = metadata[0].original_url
    
    let queueDownloads: QueueDownloads = {DownloadMedia: [request]};
    return queueDownloads
  }



  //Toast Messages
  showMessage(message: string, severity: string, summary: string) {
    this.messageService.add({ severity: severity, summary: summary, detail: message });
  }
}


