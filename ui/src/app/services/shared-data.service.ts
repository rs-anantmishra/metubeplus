import { Injectable } from '@angular/core';
import { VideoData } from '../classes/video-data';
import { BehaviorSubject, Observable, Subscription } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class SharedDataService {

    constructor() { }
    lstVideos: VideoData[] = [];
    isDownloadActive: boolean = false;
    queuedItemsMetadata: VideoData[] = [];
    activeDownloadMetadata: VideoData[] = [];

    //add or remove from queuedItems
    setQueuedItemsMetadata(metadata: VideoData[], ops: Operation) {
        if (Operation.Insert == ops) {
            let value = this.getQueuedItemsMetadata();
            value.push(...metadata)
            localStorage.setItem('queuedItemsMetadata', JSON.stringify(value));
            this.queuedItemsMetadata = value
        } else if (Operation.RemoveAtIndexZero == ops) {
            let deleted = metadata.splice(0, 1)
            localStorage.setItem('queuedItemsMetadata', JSON.stringify(metadata));
            this.queuedItemsMetadata = metadata
        }else if (Operation.Replace == ops) {
            localStorage.setItem('queuedItemsMetadata', JSON.stringify(metadata));
        }
    }

    //getQueuedItems
    getQueuedItemsMetadata() {
        let stringResult = localStorage.getItem('queuedItemsMetadata') !== null ? localStorage.getItem('queuedItemsMetadata') : JSON.stringify([])
        let queuedItemsMeta = stringResult === null ? [new VideoData()] : JSON.parse(stringResult);
        this.queuedItemsMetadata = queuedItemsMeta;

        return this.queuedItemsMetadata
    }

    setIsDownloadActive(value: boolean) {
        localStorage.setItem('isDownloadActive', JSON.stringify(value));
    }
    
    getIsDownloadActive(): boolean {
        let stringResult = localStorage.getItem('isDownloadActive') !== null ? localStorage.getItem('isDownloadActive') : 'false'
        let isActive = stringResult === null ? false : JSON.parse(stringResult);
        this.isDownloadActive = isActive;

        return this.isDownloadActive
    }

    setActiveDownloadMetadata(value: any) {
        localStorage.setItem('activeDownloadMetadata', JSON.stringify(value));
    }

    getActiveDownloadMetadata() {
        let stringResult = localStorage.getItem('activeDownloadMetadata') !== null ? localStorage.getItem('activeDownloadMetadata') : JSON.stringify([])
        let activeDownloadMeta = stringResult === null ? [new VideoData()] : JSON.parse(stringResult);
        this.activeDownloadMetadata = activeDownloadMeta;

        return this.activeDownloadMetadata
    }

    setlstVideos(value: any) {
        localStorage.setItem('lstVideos', JSON.stringify(value));
    }

    getlstVideos() {
        let stringResult = localStorage.getItem('lstVideos') !== null ? localStorage.getItem('lstVideos') : JSON.stringify([])
        let lstVideosData = stringResult === null ? [new VideoData()] : JSON.parse(stringResult);
        this.lstVideos = lstVideosData;

        return this.lstVideos
    }

    private playVideo: BehaviorSubject<VideoData> = new BehaviorSubject(new VideoData());
    onPlayVideoChange(): Observable<VideoData> {
        return this.playVideo.asObservable();
      }
      
      setPlayVideo(nextSuggestion: VideoData): void {
        this.playVideo.next(nextSuggestion);
      }
      
      resetPlayVideo(): void {
        this.playVideo.next(new VideoData());
      }
}


export enum Operation {
    Insert = 1,
    RemoveAtIndexZero = 2,
    Replace = 3
}