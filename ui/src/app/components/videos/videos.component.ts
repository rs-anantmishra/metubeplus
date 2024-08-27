import { Component, OnInit } from '@angular/core';
import { SimplecardComponent } from "../simplecard/simplecard.component";

//services
import { SharedDataService } from '../../services/shared-data.service'
import { VideosService } from '../../services/videos.service'
import { CommonModule } from '@angular/common';

@Component({
    selector: 'app-videos',
    standalone: true,
    imports: [SimplecardComponent, CommonModule],
    templateUrl: './videos.component.html',
    styleUrl: './videos.component.scss'
})

export class VideosComponent implements OnInit {

    constructor(private svcVideos: VideosService) { }

    lstVideos: any
    ngOnInit(): void {
        this.getAllVideos();
    }

    async getAllVideos() {
        let result = await this.svcVideos.getAllVideos();
        this.lstVideos = result
    }
}
