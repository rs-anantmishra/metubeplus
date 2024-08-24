import { Injectable } from '@angular/core';
import { VideoData } from '../classes/video-data';

@Injectable({
    providedIn: 'root'
})
export class SharedDataService {

    constructor() {

    }

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

}


export enum Operation {
    Insert = 1,
    RemoveAtIndexZero = 2,
    Replace = 3
}