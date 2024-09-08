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
        if (this.meta.thumbnail == '') {
            this.meta.thumbnail = './noimage.png'
        }
    }

    selectedVideo(playVideo: VideoData) {
        playVideo.video_filepath = playVideo.video_filepath.replace(/\\/g, "/");
        playVideo.video_filepath = playVideo.video_filepath.replace('../files', 'http://localhost:3500')
        // playVideo.video_filepath = playVideo.video_filepath.replace('../files', 'http://192.168.1.10:8484')
        this.svcSharedData.setPlayVideo(playVideo);
        this.router.navigate(['/play'])
    }

}
