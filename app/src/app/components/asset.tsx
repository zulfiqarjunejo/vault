"use client";

import React from "react";
import { AssetProps } from "@/app/types";
import { ChevronRightIcon } from "@heroicons/react/24/outline";
import { useRouter } from "next/navigation";

const Asset: React.FC<AssetProps> = ({ createdAt, id, name, type }) => {
  const router = useRouter();

  const handleOpenDetails = (id: string) => {
    // For example, you could navigate to a separate page to display asset details:
    router.push(`/asset-detail/${id}`);
  };

  return (
    <tr>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        {createdAt}
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        {name}
      </td>
      <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
        {type}
      </td>

      <td className="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
        <button
          className="text-blue-600 hover:text-blue-900"
          onClick={() => handleOpenDetails(id)}
        >
          <ChevronRightIcon className="h-5 w-5" aria-hidden="true" />
        </button>
      </td>
    </tr>
  );
};

export default Asset;
