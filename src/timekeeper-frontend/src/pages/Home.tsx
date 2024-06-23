import React, { useState, useEffect } from "react";
import DropdownMenu from "../components/DropdownMenu";
import VersionTable from "../components/VersionTable";
import { getAllUniqueRemoteNames, getAllRemotes, RemoteResponse } from "../api";

const Home: React.FC = () => {
  const [remoteNames, setRemoteNames] = useState<string[]>([]);
  const [selectedRemoteName, setSelectedRemoteName] = React.useState("");
  const [matchedRemotes, setMatchedRemotes] = useState<RemoteResponse[]>([]);

  useEffect(() => {
    const fetchRemoteNames = async () => {
      try {
        const names = await getAllUniqueRemoteNames();
        if (names != null) setRemoteNames(names);
      } catch (error) {
        console.error("Error fetching remote names:", error);
      }
    };

    fetchRemoteNames();
  }, []);

  useEffect(() => {
    const fetchRemotes = async () => {
      if (selectedRemoteName) {
        try {
          const remotes = await getAllRemotes(selectedRemoteName); // Replace '1.0' with the desired version
          setMatchedRemotes(remotes);
        } catch (error) {
          console.error("Error fetching remotes:", error);
        }
      }
    };

    fetchRemotes();
  }, [selectedRemoteName]);

  return (
    <div className="flex items-center justify-center h-screen bg-white px-4 sm:px-6 lg:px-8">
      <div className="max-w-[1200px] w-full flex flex-col gap-4">
        <DropdownMenu
          remoteNames={remoteNames}
          selectedRemoteName={selectedRemoteName}
          setSelectedRemoteName={setSelectedRemoteName}
        />

        <VersionTable remotes={matchedRemotes} />
      </div>
    </div>
  );
};

export default Home;
