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
      <div
        className="absolute inset-0 bg-gray-500 opacity-50"
        onClick={onClose}
      ></div>

      <div className="bg-white p-4 rounded shadow-lg relative z-50">
        <h3 className="text-lg font-semibold mb-4 text-black">Create Asset</h3>
        <div className="mb-4">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-900 mb-2"
          >
            Name
          </label>
          <input
            type="text"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="My Asset"
            className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
            required
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="name"
            className="block text-sm font-medium text-gray-900 mb-2"
          >
            Type
          </label>
          <select
            id="type"
            value={type}
            onChange={(e) => setType(e.target.value as AssetType)}
            className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
          >
            {Object.values(AssetType).map((assetType) => (
              <option key={assetType} value={assetType}>
                {assetType}
              </option>
            ))}
          </select>
        </div>

        <div className="flex justify-end space-x-4">
          <button
            onClick={handleCreateAsset}
            className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
          >
            Create
          </button>
          <button
            onClick={onClose}
            className="text-white bg-red-700 hover:bg-red-800 focus:ring-4 focus:ring-red-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
          >
            Cancel
          </button>
        </div>
      </div>
    </div>
  );
};

export default CreateAssetModal;
