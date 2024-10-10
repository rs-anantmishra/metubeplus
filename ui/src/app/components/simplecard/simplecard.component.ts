import { Component, Input, OnInit } from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { VideoData } from '../../classes/video-data'
import { CardModule } from 'primeng/card';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';
import { SharedDataService } from '../../services/shared-data.service';
import { TooltipModule } from 'primeng/tooltip';
import { TagModule } from 'primeng/tag';
import { RemovePrefixPipe } from '../../utilities/filesize-conversion.pipe'
import { CommaSepStringFromArray } from '../../utilities/array-comma-sep.pipe'
import { MinifiedViewCount } from '../../utilities/views-conversion.pipe';

@Component({
    selector: 'app-simplecard',
    standalone: true,
    imports: [ToastModule, CardModule, CommonModule, TooltipModule, TagModule, RemovePrefixPipe, CommaSepStringFromArray, MinifiedViewCount],
    providers: [MessageService, Router],
    templateUrl: './simplecard.component.html',
    styleUrl: './simplecard.component.scss'
})
export class SimplecardComponent implements OnInit {

    constructor(private router: Router, private svcSharedData: SharedDataService) {
    }
    @Input() meta: VideoData = new VideoData();

    ngOnInit(): void {
        if (this.meta.thumbnail == '') {
            this.meta.thumbnail = './noimage.png'
        }
    }

    selectedVideo(playVideo: VideoData) {
        playVideo.media_url = playVideo.media_url.replace(/\\/g, "/");
        playVideo.media_url = playVideo.media_url.replace('http://localhost:3000', 'http://localhost:3500')
        playVideo.media_url = playVideo.media_url.replace('#', '%23')

        // playVideo.media_url = playVideo.media_url.replace('../files', 'http://192.168.1.10:8484')
        this.svcSharedData.setPlayVideo(playVideo);
        this.router.navigate(['/play'])
    }

    getFormattedDuration(duration: number) {
        // Calculate minutes and seconds
        const minutes = Math.floor(duration / 60);
        const seconds = duration % 60;

        // Format the result as MM:SS
        return `${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
    }

    getFormattedUploadDate(value: string) {
        //recieved as: yyyymmdd
        let space = ' '
        let comma = ','
        const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
        let year = value.slice(0,4)
        let month = parseInt(value.slice(4,6), 10)

        let result = 'Uploaded' + space +  months[month] + comma + space + year
        return result
    }

}
