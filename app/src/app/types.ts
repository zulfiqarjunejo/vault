export enum AssetType {
    Credentials = "credentials",
    File = "file",
};

export interface AssetProps {
    createdAt: string;
    id: string;
    name: string;
    type: AssetType;
};

export interface FileUpload {
    createdAt: string;
    id: string;
    name: string;
};

export interface FileUploadListProps {
    fileUploads: FileUpload[];
};

export type CreateAssetRequest = {
    name: string;
    type: AssetType;
};

export type FileAttachmentProps = {
    assetId: string;
};
