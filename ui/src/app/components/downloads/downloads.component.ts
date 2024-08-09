
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
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { webSocket } from 'rxjs/webSocket'

//Services & Classes
import { VideoData } from '../../classes/video-data';
import { CardsService } from '../../services/cards.service';
import { SimplecardComponent } from "../simplecard/simplecard.component";

interface ExtractionOptions {
  Identifier: string;
  GetAudioOnly: boolean;
  GetSubs: boolean;
}


@Component({
  selector: 'app-downloads',
  standalone: true,
  imports: [ToastModule, ProgressBarModule, SidebarModule, CardModule, FormsModule, InputGroupModule, InputGroupAddonModule, InputTextModule, ButtonModule, CommonModule, CheckboxModule, PanelModule, SimplecardComponent],
  providers: [CardsService, MessageService],
  templateUrl: './downloads.component.html',
  styleUrl: './downloads.component.scss'
})
export class DownloadsComponent implements OnInit {

  sock = webSocket('ws://localhost:3000/ws/downloadstatus')
  msg = 'No active downloads'
  logs = ''
  downloadingChannel = 'test channel'
  downloadingTitle = 'test title'
  sendRequest() {
    this.sock.subscribe();
    this.sock.next(this.msg);
    // this.sock.complete();
    this.sock.subscribe({
      next: msg => this.updateLogs(JSON.stringify(msg)), 
      error: err => {console.log('error in ws connection', err), this.updateLogs('{"download": "ws connection closed"}')}, 
      // complete: () => console.log('complete') 
     });
  }

  updateLogs(message: string) {
    this.logs = message
    const serverlog = JSON.parse(message)
    this.logs = serverlog.download

    //these below will be set from metadata only
    //
    // this.downloadingChannel = serverlog.channel
    // this.downloadingTitle = serverlog.title
  }


  constructor(private messageService: MessageService,
    private currentDL: CardsService) { }

  interval: any;
  dlProgress: number = 98
  dl: VideoData = new VideoData()
  homeBoxActive = 'home-box'
  panelBoxActive = 'panel-box'
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
    this.sendRequest();
  }

  ngOnDestroy() {
  }



  GetMedia(): void {

    this.dl = this.currentDL.getDownloadingVideo()
    console.log(this.dl)
  }
}


