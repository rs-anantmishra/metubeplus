import { Component, OnInit } from '@angular/core';
import { SimplecardComponent } from "../simplecard/simplecard.component";

//services
import { SharedDataService } from '../../services/shared-data.service'
import { VideosService } from '../../services/videos.service'
import { CommonModule } from '@angular/common';
import { VideoData } from '../../classes/video-data';

@Component({
    selector: 'app-videos',
    standalone: true,
    imports: [SimplecardComponent, CommonModule],
    templateUrl: './videos.component.html',
    styleUrl: './videos.component.scss'
})

export class VideosComponent implements OnInit {

    constructor(private svcVideos: VideosService, private svcSharedData: SharedDataService) { }

    lstVideos: any
    ngOnInit(): void {
        this.getAllVideos();
    }

    //check local storage or service call
    async getAllVideos() {
        if (this.svcSharedData.getlstVideos().length == 0) {
            let result = await this.svcVideos.getAllVideos();
            this.lstVideos = result
            this.svcSharedData.setlstVideos(this.lstVideos)
        } else {
            this.lstVideos = this.svcSharedData.getlstVideos()
        }
    }
}
