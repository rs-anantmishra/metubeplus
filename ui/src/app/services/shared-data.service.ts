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
    videosPageSizeCount: number = -1;
    activePlayerMetadata: VideoData = new VideoData();

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
        } else if (Operation.Replace == ops) {
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

    setVideosPageSizeCount(value: any) {
        localStorage.setItem('videosPageSizeCount', JSON.stringify(value));
    }

    getVideosPageSizeCount() {
        let stringResult = localStorage.getItem('videosPageSizeCount') !== null ? localStorage.getItem('videosPageSizeCount') : JSON.stringify('-1')
        let pageSizeCount = stringResult === null ? -1 : JSON.parse(stringResult);
        this.videosPageSizeCount = pageSizeCount;

        return this.videosPageSizeCount
    }

    setActivePlayerMetadata(value: any) {
        localStorage.setItem('activePlayerMetadata', JSON.stringify(value));
    }

    getActivePlayerMetadata() {
        let stringResult = localStorage.getItem('activePlayerMetadata') !== null ? localStorage.getItem('activePlayerMetadata') : JSON.stringify(new VideoData())
        let activePlayerMeta = stringResult === null ? new VideoData() : JSON.parse(stringResult);
        this.activePlayerMetadata = activePlayerMeta;

        return this.activePlayerMetadata
    }



    private playVideo: BehaviorSubject<VideoData> = new BehaviorSubject(new VideoData());
    onPlayVideoChange(): Observable<VideoData> {        
        //check localstorage
        let activeVideo = this.getActivePlayerMetadata()
        if (activeVideo.video_filepath != '') {
            this.setPlayVideo(activeVideo)
        }
        return this.playVideo.asObservable();
    }

    setPlayVideo(data: VideoData): void {
        this.setActivePlayerMetadata(data);
        this.playVideo.next(data);
    }

    resetPlayVideo(): void {
        this.playVideo.next(new VideoData());
    }

    private pageSizeCount: BehaviorSubject<number> = new BehaviorSubject(-1)
    setPageSizeCount(count: number): void {
        this.setVideosPageSizeCount(count);
        this.pageSizeCount.next(count);
    }

    getPageSizeCount(): Observable<number> {
        return this.pageSizeCount.asObservable()
    }
}


export enum Operation {
    Insert = 1,
    RemoveAtIndexZero = 2,
    Replace = 3
}