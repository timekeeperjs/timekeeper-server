import React from "react";
import { FaTimeline } from "react-icons/fa6";
import { RemoteResponse } from "../api";

interface IVersionTable {
  remotes: RemoteResponse[];
}

const VersionTable: React.FC<IVersionTable> = ({ remotes }) => {
  return (
    <div className="overflow-hidden rounded-lg border border-gray-200 shadow-md">
      <table className="min-w-full divide-y divide-gray-200">
        <thead className="bg-gray-50">
          <tr>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Remote Name
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Version
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Remote URL
            </th>
            <th
              scope="col"
              className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
            >
              Entrypoint
            </th>
          </tr>
        </thead>
        <tbody className="bg-white divide-y divide-gray-200">
          {remotes.length === 0 ? (
            <tr>
              <td
                colSpan={4}
                className="px-6 py-4 whitespace-nowrap text-center"
              >
                <div className="flex flex-col items-center justify-center">
                  <FaTimeline className="text-blue-500 text-6xl mb-4" />
                  <p className="text-lg font-medium text-gray-900">
                    Select a property by searching for a hotel name or ID.
                  </p>
                </div>
              </td>
            </tr>
          ) : (
            remotes.map((remote, index) => (
              <tr key={index}>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {remote.remoteName}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {remote.version}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {remote.remoteURL}
                </td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {remote.remoteURL}
                  {remote.version}.remoteEntry.js
                </td>
              </tr>
            ))
          )}
        </tbody>
      </table>
    </div>
  );
};

export default VersionTable;
