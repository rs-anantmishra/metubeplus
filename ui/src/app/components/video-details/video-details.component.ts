import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SharedDataService } from '../../services/shared-data.service';
import { VideoData } from '../../classes/video-data';
import Plyr from 'plyr';


@Component({
    selector: 'app-video-details',
    standalone: true,
    imports: [CommonModule, RouterModule, ButtonModule],
    providers: [Router],
    templateUrl: './video-details.component.html',
    styleUrl: './video-details.component.scss'
})
export class VideoDetailsComponent implements OnInit {

    player: any;
    selectedVideo: VideoData = new VideoData()

    constructor(private svcSharedData: SharedDataService) { 
        
    }
    ngOnInit(): void {        
        this.svcSharedData.onPlayVideoChange().subscribe(selectedVideo => this.selectedVideo = selectedVideo);
        this.player = new Plyr('#plyrID', { captions: { active: true } });
    }
}
