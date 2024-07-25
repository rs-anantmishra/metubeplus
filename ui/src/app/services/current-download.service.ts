import { Injectable } from '@angular/core';
import { VideoData } from '../classes/video-data';

@Injectable({
  providedIn: 'root'
})
export class CurrentDownloadService {

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
  constructor() { }

  getDownloaingVideo(): VideoData {
    let data = this.videoData.find((x) => x.Title === 'Test Title')
    return data!;
  }
}


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