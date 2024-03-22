import React, { useState } from 'react';
import config from '@/app/config';
import { FileAttachmentProps } from '@/app/types';

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
      formData.append('file', selectedFile);

      const response = await fetch(`${config.apiUrl}/assets/${assetId}/files`, {
        method: 'POST',
        body: formData,
      });

      if (!response.ok) {
        throw new Error('Failed to upload file');
      }
      
      // Optionally, you can perform additional actions after successful upload
    } catch (error) {
      console.error('Error uploading file:', error);
      // Optionally, you can handle and display the error to the user
    }
  };

  return (
    <>
      <label htmlFor="file" className="mb-[10px] block text-base font-medium text-dark dark:text-white">
        Attachments
      </label>
      <div className="relative">
        <label
          htmlFor="file"
          className="flex min-h-[175px] w-full cursor-pointer items-center justify-center rounded-md border border-dashed border-primary p-6"
        >
          <div>
            <input 
              type="file" 
              name="file" 
              id="file" 
              className="sr-only" 
              onChange={handleFileChange} 
            />
            <span className="mx-auto mb-3 flex h-[50px] w-[50px] items-center justify-center rounded-full border border-stroke dark:border-dark-3 bg-white dark:bg-dark-2">
              <svg
                width="20"
                height="20"
                viewBox="0 0 20 20"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                {/* SVG icon for file upload */}
              </svg>
            </span>
            <span className="text-base text-body-color dark:text-dark-6">
              Drag &amp; drop or
              <span className="text-primary underline"> browse </span>
            </span>
          </div>
        </label>
      </div>
      <button 
        type="button" 
        onClick={handleUploadFile} 
        disabled={!selectedFile} 
        className="bg-primary text-white px-3 py-2 rounded-md mt-4 hover:bg-primary-dark focus:outline-none focus:ring-2 focus:ring-primary"
      >
        Upload File
      </button>
    </>
  );
};

export default FileAttachment;
