import { Title } from "@angular/platform-browser"

export class VideoData {
    video_id: number = -1
    title: string = ''
    description: string = ''
    duration: number = -1
    original_url: string = ''
    is_file_downloaded: boolean = false
    channel: string = ''
    playlist: string = ''
    playlist_video_index: number = -1
    domain: string = ''
    video_format: string = ''
    watch_count: number = -1
    is_deleted: boolean = false
    created_date: number = -1
    thumbnail: string = ''
}

export class VideoDataRequest {
    Indicator: string = ''
    SubtitlesReq: boolean = false
    IsAudioOnly: boolean = false
}

export class QueueDownloads {
    DownloadMedia: DownloadMedia[] = [new DownloadMedia()]
}

export class DownloadMedia {
    VideoId: number = -1
    VideoURL: string = ''
}