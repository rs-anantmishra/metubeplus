import { Title } from "@angular/platform-browser"

export class VideoData {
    Title: string = ''
    Description: string = ''
    Duration: number = -1
    OriginalURL: string = ''
    IsFileDownloaded: boolean = false
    Channel: string = ''
    Playlist: string = ''
    PlaylistVideoIndex: number = -1
    Domain: string = ''
    VideoFormat: string = ''
    WatchCount: number = -1
    IsDeleted: boolean = false
    CreatedDate: number = -1
}

export class VideoDataRequest {
    Indicator: string = ''
    SubtitlesReq: boolean = false
    IsAudioOnly: boolean = false
}