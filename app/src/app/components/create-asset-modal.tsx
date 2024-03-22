// components/CreateAssetModal.tsx

import React, { useState } from "react";
import { CreateAssetRequest, AssetType } from "@/app/types";

const CreateAssetModal: React.FC<{
  isOpen: boolean;
  onClose: () => void;
  onCreateAsset: (asset: CreateAssetRequest) => void;
}> = ({ isOpen, onClose, onCreateAsset }) => {
  const [name, setName] = useState("");
  const [type, setType] = useState(AssetType.Credentials);

  const handleCreateAsset = () => {
    // Check if name is not empty
    if (name.trim() === "") {
      // Optionally, you can show an error message to the user
      return;
    }

    // Call the onCreateAsset callback with the new asset object
    onCreateAsset({ name, type });

    // Clear input fields
    setName("");

    // Close the modal
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 flex items-center justify-center z-50">
      {" "}
      {/* Increase z-index */}
      {/* Overlay */}
      <div
        className="absolute inset-0 bg-gray-500 opacity-50"
        onClick={onClose}
      ></div>
      {/* Modal */}
      <div className="bg-white p-4 rounded shadow-lg relative z-50">
        {" "}
        {/* Increase z-index */}
        <h3 className="text-lg font-semibold mb-4">Create Asset</h3>
        <div className="mb-4">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-700"
          >
            Name
          </label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full text-black"
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="type"
            className="block text-sm font-medium text-gray-700"
          >
            Type
          </label>
          <select
            id="type"
            value={type}
            onChange={(e) => setType(e.target.value as AssetType)}
            className="mt-1 p-2 border border-gray-300 rounded-md w-full text-black"
          >
            {Object.values(AssetType).map((assetType) => (
              <option key={assetType} value={assetType}>
                {assetType}
              </option>
            ))}
          </select>
        </div>
        <div className="flex justify-end">
          <button
            onClick={handleCreateAsset}
            className="px-3 py-1 bg-indigo-500 text-white rounded hover:bg-indigo-600 focus:outline-none focus:bg-indigo-600"
          >
            Create
          </button>
          <button
            onClick={onClose}
            className="ml-2 px-3 py-1 border border-red-500 rounded text-red-500 hover:bg-red-100 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
};

export default CreateAssetModal;
