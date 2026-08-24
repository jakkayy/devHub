'use client';

import React, { useState } from 'react';
import { ExternalLink, Maximize2, Minimize2, Network } from 'lucide-react';

interface MiroWidgetProps {
  boardId?: string;
}

export default function MiroWidget({ boardId }: MiroWidgetProps) {
  const [isFullscreen, setIsFullscreen] = useState(false);

  // Valid embed URL for Miro Board
  const embedUrl = boardId
    ? `https://miro.com/app/live-embed/${boardId}/?embedMode=view_only_without_ui`
    : 'https://miro.com/app/embed/uXjV.../?embedMode=view_only_without_ui';

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
              className="p-1.5 rounded-lg bg-zinc-800 hover:bg-zinc-700 text-zinc-300 transition flex items-center gap-1 text-xs"
              title="Fullscreen"
            >
              <Maximize2 className="w-3.5 h-3.5" /> Fullscreen
            </button>
            <span className="text-xs text-zinc-400">Miro Viewer</span>
          </div>
        </div>

        {/* Miro Placeholder & Embedded Interactive Viewer Container */}
        <div className="relative flex-1 min-h-[300px] w-full rounded-lg overflow-hidden border border-zinc-800 bg-zinc-950 flex flex-col items-center justify-center p-6 text-center">
          <div className="p-4 rounded-full bg-amber-500/10 border border-amber-500/20 text-amber-400 mb-3">
            <Network className="w-8 h-8" />
          </div>
          <h3 className="text-base font-semibold text-zinc-200">System Architecture & Workflow Board</h3>
          <p className="text-xs text-zinc-400 max-w-md mt-1 mb-4">
            Interactive Miro board embedded view for system design, component dependencies, and node-to-code mapping.
          </p>
          <a
            href="https://miro.com"
            target="_blank"
            rel="noreferrer"
            className="inline-flex items-center gap-1.5 text-xs text-amber-400 hover:text-amber-300 font-medium bg-amber-950/60 border border-amber-800/40 px-3 py-1.5 rounded-md transition"
          >
            Open Live Board in Miro <ExternalLink className="w-3.5 h-3.5" />
          </a>
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
                href="https://miro.com"
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

          <div className="flex-1 w-full rounded-xl overflow-hidden border border-zinc-800 bg-zinc-950 flex flex-col items-center justify-center p-8 text-center">
            <Network className="w-12 h-12 text-amber-400 mb-4" />
            <h3 className="text-lg font-bold text-zinc-100">Miro Architecture Board Viewer</h3>
            <p className="text-sm text-zinc-400 max-w-lg mt-1 mb-6">
              View system design nodes and map them directly with code repository files.
            </p>
            <a
              href="https://miro.com"
              target="_blank"
              rel="noreferrer"
              className="inline-flex items-center gap-2 text-sm text-amber-400 hover:text-amber-300 font-medium bg-amber-950/60 border border-amber-800/40 px-4 py-2 rounded-lg transition"
            >
              Open Full Live Board in Miro <ExternalLink className="w-4 h-4" />
            </a>
          </div>
        </div>
      )}
    </>
  );
}
