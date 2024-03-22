import React from "react";
import { FileUploadListProps } from "@/app/types";

const FileUploadList: React.FC<FileUploadListProps> = ({ fileUploads }) => {
  return (
    <div>
      <h3 className="text-lg font-medium mb-2">File Uploads</h3>
      <table className="min-w-full">
        <thead>
          <tr>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
              Created
            </th>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"></th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {fileUploads.map((upload) => (
            <tr key={upload.id}>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {upload.createdAt}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {upload.name}
              </td>
              <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500"></td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default FileUploadList;
