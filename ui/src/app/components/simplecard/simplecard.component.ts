import { Component, Input } from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { MessageService } from 'primeng/api';
import { PanelModule } from 'primeng/panel';
import { VideoData } from '../../classes/video-data'
import { CardModule } from 'primeng/card';

@Component({
    selector: 'app-simplecard',
    standalone: true,
    imports: [ToastModule, PanelModule, CardModule],
    providers: [MessageService],
    templateUrl: './simplecard.component.html',
    styleUrl: './simplecard.component.scss'
})
export class SimplecardComponent {

    @Input() meta: VideoData = new VideoData();


    getChannel() {
        if (this.meta.channel === '') {
            return 'No Downloads Queued!'
        } else {
            return this.meta.channel;
        }
    }
}
