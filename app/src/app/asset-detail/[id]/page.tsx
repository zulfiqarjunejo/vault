"use client";

import React, { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { AssetProps, FileUpload } from "@/app/types";
import config from "@/app/config";
import FileUploadList from "@/app/components/file-upload-list";
import AssetInformation from "@/app/components/asset-information";
import BackButton from "@/app/components/back-button";
import FileAttachment from "@/app/components/file-attachment";

const AssetDetail: React.FC = () => {
  const { id } = useParams();

  const [asset, setAsset] = useState<AssetProps | null>(null);
  const [fileUploads, setFileUploads] = useState<FileUpload[]>([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        // Fetch asset details
        const assetResponse = await fetch(`${config.apiUrl}/assets/${id}`);
        if (!assetResponse.ok) {
          throw new Error("Failed to fetch asset details");
        }
        const assetData: AssetProps = await assetResponse.json();
        setAsset(assetData);

        // Fetch file uploads associated with the asset
        const fileUploadResponse = await fetch(
          `${config.apiUrl}/assets/${id}/files`
        );
        if (!fileUploadResponse.ok) {
          throw new Error("Failed to fetch file uploads");
        }
        const fileUploadsData: FileUpload[] = await fileUploadResponse.json();
        setFileUploads(fileUploadsData);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    if (id) {
      fetchData();
    }
  }, [id]);

  return (
    <div className="container mx-auto p-4">
      <h2 className="text-2xl font-bold mb-4 flex items-center justify-between">
        Asset Details
        <BackButton />
      </h2>
      {asset && <AssetInformation asset={asset} />}

      {asset && <FileAttachment assetId={asset!.id} />}

      {fileUploads?.length > 0 && <FileUploadList fileUploads={fileUploads} />}
    </div>
  );
};

export default AssetDetail;
