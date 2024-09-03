import { Component, OnDestroy, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Router, RouterModule } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SharedDataService } from '../../services/shared-data.service';
import { VideoData } from '../../classes/video-data';
import { Subscription } from 'rxjs';
import { ScrollPanelModule } from 'primeng/scrollpanel';
import Plyr from 'plyr';


@Component({
    selector: 'app-video-details',
    standalone: true,
    imports: [CommonModule, RouterModule, ButtonModule, ScrollPanelModule],
    providers: [Router],
    templateUrl: './video-details.component.html',
    styleUrl: './video-details.component.scss'
})
export class VideoDetailsComponent implements OnInit, OnDestroy {

    subscription!: Subscription;
    player: any;
    selectedVideo: VideoData = new VideoData()

    constructor(private svcSharedData: SharedDataService) {
        this.subscription = this.svcSharedData.onPlayVideoChange().subscribe(selectedVideo => this.selectedVideo = selectedVideo);
        this.player = new Plyr('#plyrID', { captions: { active: true }, loop: { active: true }, ratio: '16:9' });
    }

    async ngOnInit(): Promise<void> {
    }

    ngOnDestroy(): void {
        this.subscription.unsubscribe()
    }
}
