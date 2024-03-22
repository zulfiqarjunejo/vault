import Asset from "@/app/components/asset";
import CreateAssetModal from "@/app/components/create-asset-modal";
import { AssetProps, CreateAssetRequest } from "@/app/types";
import { PlusIcon } from "@heroicons/react/24/outline";
import React, { useEffect, useState } from "react";
import config from "../config";

const AssetList: React.FC = () => {
  const [assets, setAssets] = useState<AssetProps[]>([]);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const handleOpenModal = () => {
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    setIsModalOpen(false);
  };

  const handleCreateAsset = async (createAssetRequest: CreateAssetRequest) => {
    const response = await fetch(`${config.apiUrl}/assets`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(createAssetRequest),
    });

    if (!response.ok) {
      throw new Error("Failed to create asset");
    }

    fetchAssets();
  };

  useEffect(() => {
    fetchAssets();
  }, []);

  const fetchAssets = async () => {
    try {
      const response = await fetch(`${config.apiUrl}/assets`);
      if (!response.ok) {
        throw new Error("Failed to fetch assets");
      }

      const data: AssetProps[] = await response.json();
      setAssets(data);
    } catch (error) {
      console.error("Error fetching assets:", error);
    }
  };

  const handleAddAsset = () => {};

  return (
    <div className="container mx-auto p-4">
      <h2 className="text-2xl font-bold mb-4 flex items-center justify-between">
        Asset List
        <button
          type="button"
          onClick={handleOpenModal}
          className="inline-flex items-center justify-center w-8 h-8 border border-gray-300 rounded-full bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          <PlusIcon className="w-4 h-4 text-gray-700" />
        </button>
        <CreateAssetModal
          isOpen={isModalOpen}
          onClose={handleCloseModal}
          onCreateAsset={handleCreateAsset}
        />
      </h2>

      <table className="min-w-full">
        <thead>
          <tr>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
              Created
            </th>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
              Name
            </th>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider">
              Type
            </th>
            <th className="px-6 py-3 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"></th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {assets?.map((asset) => (
            <Asset key={asset.id} {...asset} />
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default AssetList;
