export interface DownloadItem {
    link: string
    size: number
    checksumLink: string
}

export interface PatchItem {
    fromBuild: string
    link: string
    size: number
    checksumLink: string
}

export interface ReleaseData {
    date: string
    type: string
    downloads: {
        [key: string]: DownloadItem
    }
    patches?: {
        win?: PatchItem[]
        mac?: PatchItem[]
        unix?: PatchItem[]
    }
    notesLink?: string
    licenseRequired?: boolean
    version: string
    majorVersion: string
    build: string
    whatsnew?: string
    uninstallFeedbackLinks?: {
        [key: string]: string
    }
    printableReleaseType?: string | null
}

export interface DownloadData {
    [key: string]: ReleaseData[]
}

export interface DownloadOption {
    name: string
    size: string
    format: string
    link: string
    recommended: boolean
}