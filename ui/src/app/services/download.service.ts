import { Injectable } from '@angular/core';
import { VideoData, VideoDataRequest, QueueDownloads } from '../classes/video-data';
import { HttpClient } from '@angular/common/http';

const apiUrl: string = 'http://localhost:3000/api'

@Injectable({
  providedIn: 'root'
})
export class DownloadService {

  constructor(private http: HttpClient) { }

  getDownloadingVideo(): string {
    return '';
  }

  //metadata
  async getMetadata(request: VideoDataRequest): Promise<VideoData[]> {
    let url = '/download/metadata'

    return fetch(apiUrl + url, {
      method: 'POST',
      body: JSON.stringify(request),
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => { return response.json(); });
  }

  //media
  async getMedia(request: QueueDownloads): Promise<string> {
    let url = '/download/media'

    return fetch(apiUrl + url, {
      method: 'POST',
      body: JSON.stringify(request),
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => { return response.json(); });
  }

  
}