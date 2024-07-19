
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { InputGroupModule } from 'primeng/inputgroup';
import { InputGroupAddonModule } from 'primeng/inputgroupaddon';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';
import { CheckboxModule } from 'primeng/checkbox';
import { PanelModule } from 'primeng/panel';

interface ExtractionOptions {
  Identifier: string;
  GetAudioOnly: boolean;
  GetSubs: boolean;
}


@Component({
  selector: 'app-downloads',
  standalone: true,
  imports: [FormsModule, InputGroupModule, InputGroupAddonModule, InputTextModule, ButtonModule, CommonModule, CheckboxModule, PanelModule],
  templateUrl: './downloads.component.html',
  styleUrl: './downloads.component.scss'
})
export class DownloadsComponent implements OnInit {
  homeBoxActive = 'home-box'
  panelBoxActive = 'panel-box'

  options :ExtractionOptions = { Identifier: '', GetAudioOnly: false, GetSubs: false}

  flipCheckbox(event: any, option: string): void {
    if(option === 'GetAudioOnly') {
      this.options.GetAudioOnly = !event.checked;
    } else if (option === 'GetSubs') {
      this.options.GetSubs = !event.checked;
    }
  }

  placeholder = 'Video or Playlist URL'
  ngOnInit(): void {

  }

  GetMedia(): void {
    this.homeBoxActive = 'home-box-queued'
    setTimeout(()=>{
      this.panelBoxActive = 'panel-box-queued'
    },100);
  }
}


