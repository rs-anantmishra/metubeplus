export class Messages {
    //arrays
    Severities: string[] = ['Secondary', 'Success', 'Info', 'Warning', 'Help', 'Danger', 'Contrast']

    //literals
    wsMessage: string = 'No active downloads.'
    serverLogs: string = 'No logs available.'
    downloadComplete: string = 'Download completed successfully.'
    triggerDownloadApiSuccessResponse: string = 'Item added to download queue successfully.'
    downloadInfoIdentifier: string = '[download]'
}

export const wsApiUrl: string = 'ws://localhost:3000/ws/downloadstatus'
export enum Severity {
    secondary = 1,
    success = 2,
    info = 3,
    warning = 4,
    help = 5,
    danger = 6,
    contrast = 7
}