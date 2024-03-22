import React, { useState } from "react";
import config from "@/app/config";
import { FileAttachmentProps } from "@/app/types";
import { FolderIcon } from "@heroicons/react/24/outline";

const FileAttachment: React.FC<FileAttachmentProps> = ({ assetId }) => {
  const [selectedFile, setSelectedFile] = useState<File | null>(null);

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      setSelectedFile(file);
    }
  };

  const handleUploadFile = async () => {
    if (!selectedFile) return;

    try {
      const formData = new FormData();
      formData.append("file", selectedFile);

      const response = await fetch(`${config.apiUrl}/assets/${assetId}/files`, {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error("Failed to upload file");
      }

      window.location.reload();
    } catch (error) {
      console.error("Error uploading file:", error);
    }
  };

  return (
    <>
      <label
        htmlFor="file"
        className="flex w-full cursor-pointer items-center justify-center rounded-md border border-dashed border-primary p-6"
      >
        <div>
          <input
            type="file"
            name="file"
            id="file"
            className="sr-only"
            onChange={handleFileChange}
          />
          <span className="mx-auto mb-3 flex h-[50px] w-[50px] items-center justify-center">
            <FolderIcon className="flex h-[50px] w-[50px]"/>
          </span>

          {selectedFile?.name ? (
            <span className="text-base text-body-color dark:text-dark-6">
              {selectedFile.name}
            </span>
          ) : (
            <span className="text-base text-body-color dark:text-dark-6">
              Drag &amp; drop or
              <span className="text-primary underline"> browse </span>
            </span>
          )}
        </div>
      </label>

      <button
        type="button"
        onClick={handleUploadFile}
        disabled={!selectedFile}
        className="bg-primary text-white px-3 py-2 rounded-md mt-4 outline hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-primary"
      >
        Upload File
      </button>
    </>
  );
};

export default FileAttachment;
