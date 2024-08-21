import { Injectable } from '@angular/core';
import { VideoData, VideoDataRequest } from '../classes/video-data';
import { HttpClient } from '@angular/common/http';

const server: string = 'http://localhost:3000'

@Injectable({
  providedIn: 'root'
})
export class DownloadService {

  protected videoData: VideoData[] = [
    {
      Title: 'Demo Title',
      Description: 'Demo Description',
      Duration: 100,
      OriginalURL: 'http://demo.youtube.com',
      IsFileDownloaded: false,
      Channel: 'Demo Channel',
      Playlist: '',
      PlaylistVideoIndex: -1,
      Domain: 'youtube.com',
      VideoFormat: 'Demo Format',
      WatchCount: -1,
      IsDeleted: false,
      CreatedDate: 1
    },
    {
      Title: 'Test Title',
      Description: 'Test Description',
      Duration: 100,
      OriginalURL: 'http://test.youtube.com',
      IsFileDownloaded: false,
      Channel: 'Test Channel',
      Playlist: '',
      PlaylistVideoIndex: -1,
      Domain: 'youtube.com',
      VideoFormat: 'Test Format',
      WatchCount: -1,
      IsDeleted: false,
      CreatedDate: 1
    }

  ]; 
  constructor(private http: HttpClient) { }

  getDownloadingVideo(): VideoData {
    let data = this.videoData.find((x) => x.Title === 'Test Title 1')
    return data!;
  }

  meta!: any
  getMetadata(request: VideoDataRequest): VideoData {
    console.log(1)
    debugger
    let metadata: any
    let url = server + '/download/metadata'
    let result = this.http.post(url, request).subscribe(res => {
      this.meta = res
    })
    
    return metadata!
  }
}

// {
//   "Indicator":"PLIhvC56v63IKrRHh3gvZZBAGvsvOhwrRF", 
//   "SubtitlesReq": false,
//   "IsAudioOnly": false
// }


// {
//   title: string = ''
//   description: string = ''
//   duration: number = -1
//   originalURL: string = ''
//   isFileDownloaded: boolean = false
//   Channel: string = ''
//   Playlist: string = ''
//   PlaylistVideoIndex: number = -1
//   Domain: string = ''
//   VideoFormat: string = ''
//   WatchCount: number = -1
//   IsDeleted: boolean = false
//   CreatedDate: number = -1
// }

// {
//   "Indicator":"PLIhvC56v63IKrRHh3gvZZBAGvsvOhwrRF", 
//   "SubtitlesReq": false,
//   "IsAudioOnly": false
// }