'use client';

import React, { useState } from 'react';
import { ExternalLink, Maximize2, Minimize2, Network } from 'lucide-react';

interface MiroWidgetProps {
  boardId?: string;
}

export default function MiroWidget({ boardId }: MiroWidgetProps) {
  const [isFullscreen, setIsFullscreen] = useState(false);

  // Default demo Miro board embed link or custom board ID
  const embedUrl = boardId
    ? `https://miro.com/app/live-embed/${boardId}/?embedMode=view_only_without_ui`
    : 'https://miro.com/app/live-embed/uXjVO1S_demo=/?embedMode=view_only_without_ui';

  return (
    <>
      <div className="bg-zinc-900/90 border border-zinc-800 rounded-xl p-5 shadow-lg flex flex-col h-full">
        <div className="flex items-center justify-between mb-4">
          <div className="flex items-center gap-2">
            <Network className="w-5 h-5 text-amber-400" />
            <h2 className="text-lg font-semibold text-zinc-100">Workflow & Architecture Diagram</h2>
          </div>
          <div className="flex items-center gap-2">
            <button
              onClick={() => setIsFullscreen(true)}
              className="p-1.5 rounded-lg bg-zinc-800 hover:bg-zinc-700 text-zinc-300 transition"
              title="Fullscreen"
            >
              <Maximize2 className="w-4 h-4" />
            </button>
            <span className="text-xs text-zinc-400">Miro Embed</span>
          </div>
        </div>

        {/* Miro iFrame Viewer Container */}
        <div className="relative flex-1 min-h-[300px] w-full rounded-lg overflow-hidden border border-zinc-800 bg-zinc-950">
          <iframe
            src={embedUrl}
            className="w-full h-full min-h-[300px] border-0"
            allow="fullscreen; clipboard-read; clipboard-write"
            title="Miro Architecture Diagram"
          ></iframe>
        </div>
      </div>

      {/* Fullscreen Modal View */}
      {isFullscreen && (
        <div className="fixed inset-0 z-50 bg-zinc-950/90 backdrop-blur flex flex-col p-6">
          <div className="flex items-center justify-between mb-4">
            <div className="flex items-center gap-2">
              <Network className="w-6 h-6 text-amber-400" />
              <h2 className="text-xl font-bold text-zinc-100">Miro Workflow Diagram (Fullscreen)</h2>
            </div>
            <div className="flex items-center gap-3">
              <a
                href={embedUrl}
                target="_blank"
                rel="noreferrer"
                className="flex items-center gap-1 text-sm text-amber-400 hover:text-amber-300 font-medium transition"
              >
                Open in Miro <ExternalLink className="w-4 h-4" />
              </a>
              <button
                onClick={() => setIsFullscreen(false)}
                className="p-2 rounded-lg bg-zinc-800 hover:bg-zinc-700 text-zinc-300 transition"
              >
                <Minimize2 className="w-5 h-5" />
              </button>
            </div>
          </div>

          <div className="flex-1 w-full rounded-xl overflow-hidden border border-zinc-800 bg-zinc-950">
            <iframe
              src={embedUrl}
              className="w-full h-full border-0"
              allow="fullscreen; clipboard-read; clipboard-write"
              title="Miro Architecture Diagram Fullscreen"
            ></iframe>
          </div>
        </div>
      )}
    </>
  );
}
