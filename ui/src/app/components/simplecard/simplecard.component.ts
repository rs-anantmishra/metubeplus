import { Component, Input, OnInit } from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { VideoData } from '../../classes/video-data'
import { CardModule } from 'primeng/card';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { SharedDataService } from '../../services/shared-data.service';

@Component({
    selector: 'app-simplecard',
    standalone: true,
    imports: [ToastModule, CardModule, CommonModule],
    providers: [MessageService, Router],
    templateUrl: './simplecard.component.html',
    styleUrl: './simplecard.component.scss'
})
export class SimplecardComponent implements OnInit {

    constructor(private router: Router, private svcSharedData: SharedDataService) {
    }
    // meta: VideoData = new VideoData()
    @Input() meta: VideoData = new VideoData();

    ngOnInit(): void {

        console.log('nav:', this.router.navigated)
        if (this.meta.thumbnail == '') {
            this.meta.thumbnail = './noimage.png'
        }
    }

    getVideoDetails(playVideo: VideoData) {
        // console.log('videoId', playVideo)
        this.svcSharedData.setPlayVideo(playVideo);
        if (!this.router.navigated) {
            this.router.navigate(['/play'])
        } else {
            console.log('navigated')
        }
    }

}
