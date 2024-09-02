import { Component, OnInit } from '@angular/core';
import { PaginatorModule } from 'primeng/paginator';
import { ButtonModule } from 'primeng/button';
import { ScrollPanelModule } from 'primeng/scrollpanel';
import { SimplecardComponent } from "../simplecard/simplecard.component";

//services
import { SharedDataService } from '../../services/shared-data.service'
import { VideosService } from '../../services/videos.service'
import { CommonModule } from '@angular/common';
import { VideoData } from '../../classes/video-data';
import { Scroll } from '@angular/router';

interface PageEvent {
    first: number;
    rows: number;
    page: number;
    pageCount: number;
}

@Component({
    selector: 'app-videos',
    standalone: true,
    imports: [SimplecardComponent, CommonModule, PaginatorModule, ButtonModule, ScrollPanelModule],
    templateUrl: './videos.component.html',
    styleUrl: './videos.component.scss'
})

export class VideosComponent implements OnInit {
    visibility = 'visible'
    first: number = 0;
    rows: number = 10;
    totalRecords: number = -1

    //videos presenter
    lstVideos: any

    constructor(private svcVideos: VideosService, private svcSharedData: SharedDataService) { }
    ngOnInit(): void {
        this.getAllVideos();
        //this.getAllVideosDelta();
    }

    //check local storage or service call
    async getAllVideosDelta() {
        if (this.svcSharedData.getlstVideos().length == 0) {
            let result = await this.svcVideos.getAllVideos();
            this.svcSharedData.setlstVideos(result)
            this.lstVideos = this.getPagedResult(this.first, this.rows);
        } else {
            this.lstVideos = this.getPagedResult(this.first, this.rows);
        }
    }

    async getAllVideos() {
        let result = await this.svcVideos.getAllVideos();
        this.svcSharedData.setlstVideos(result)
        this.lstVideos = this.getPagedResult(this.first, this.rows);
    }

    getPagedResult(first: number, rows: number): any {
        let result = this.svcSharedData.getlstVideos()
        this.totalRecords = result.length
        return result.slice(first, (first + rows))
    }

    onPageChange(event: any) {
        //set array to match page
        this.lstVideos = this.getPagedResult(event.first, event.rows)
        this.first = event.first
        this.rows = event.rows;
    }
}
