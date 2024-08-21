import { Injectable } from '@angular/core';
import { VideoData, VideoDataRequest } from '../classes/video-data';
import { HttpClient } from '@angular/common/http';

const apiUrl: string = 'http://localhost:3000'

@Injectable({
  providedIn: 'root'
})
export class DownloadService {

  constructor(private http: HttpClient) { }

  getDownloadingVideo(): string {
    return '';
  }

  async getMetadata(request: VideoDataRequest): Promise<any> {
    let url = '/download/metadata'

    return fetch(apiUrl + url, {
      method: 'POST',
      body: JSON.stringify(request),
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(response => { return response.json(); });
  }
}