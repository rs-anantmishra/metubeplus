import { Component, Input, OnInit } from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { PanelModule } from 'primeng/panel';
import { VideoData } from '../../classes/video-data'
import { CardModule } from 'primeng/card';
import { CommonModule } from '@angular/common';

@Component({
    selector: 'app-simplecard',
    standalone: true,
    imports: [ToastModule, PanelModule, CardModule, CommonModule],
    providers: [MessageService],
    templateUrl: './simplecard.component.html',
    styleUrl: './simplecard.component.scss'
})
export class SimplecardComponent implements OnInit {

    @Input() meta: VideoData = new VideoData();


    ngOnInit(): void {
        if (this.meta.thumbnail == ''){
            this.meta.thumbnail = './noimage.png'
        }
    }
}
