import React, { useState } from "react";
import { FiChevronDown } from "react-icons/fi";

interface IDropdownMenu {
  remoteNames: string[];
  selectedRemoteName: string;
  setSelectedRemoteName(remoteName: string): void;
}

const DropdownMenu: React.FC<IDropdownMenu> = ({
  remoteNames,
  selectedRemoteName,
  setSelectedRemoteName,
}) => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleDropdown = () => {
    setIsOpen(!isOpen);
  };

  const handleOptionClick = (remoteName: string) => {
    setSelectedRemoteName(remoteName);
    setIsOpen(false);
  };

  return (
    <div className="relative inline-block text-left w-64">
      <label className="block text-sm font-medium text-gray-700 mb-2">
        Select Remote Name
      </label>
      <div>
        <button
          onClick={toggleDropdown}
          className="inline-flex justify-between items-center w-full rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-sm font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          <span>{selectedRemoteName || "Select"}</span>
          <FiChevronDown className="h-5 w-5" />
        </button>
      </div>
      {isOpen && (
        <div className="origin-top-right absolute right-0 mt-2 w-full rounded-md shadow-lg bg-white ring-1 ring-black ring-opacity-5 focus:outline-none">
          <div className="py-1">
            {remoteNames.map((remoteName, index) => (
              <button
                key={index}
                onClick={() => handleOptionClick(remoteName)}
                className="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100 w-full text-left"
              >
                {remoteName}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default DropdownMenu;
