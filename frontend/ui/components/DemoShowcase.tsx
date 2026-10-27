"use client";

import React, { useEffect, useRef, useState, useCallback } from "react";
import { FaPlay } from "react-icons/fa6";

interface ClipItem {
  id: number;
  duration: string;
  image: string;
}

const CLIPS: ClipItem[] = [
  { id: 1, duration: "0:34", image: "/CLIP_1.jpg" },
  { id: 2, duration: "0:45", image: "/CLIP_2.jpg" },
  { id: 3, duration: "0:28", image: "/CLIP_3.jpg" },
  { id: 4, duration: "0:52", image: "/CLIP_4.jpg" },
];

interface LinePath {
  d: string;
  startPoint: { x: number; y: number };
  endPoint: { x: number; y: number };
}

export default function DemoShowcase() {
  const containerRef = useRef<HTMLDivElement>(null);
  const centerRef = useRef<HTMLDivElement>(null);
  const clip1Ref = useRef<HTMLDivElement>(null);
  const clip2Ref = useRef<HTMLDivElement>(null);
  const clip3Ref = useRef<HTMLDivElement>(null);
  const clip4Ref = useRef<HTMLDivElement>(null);

  const [hoveredClip, setHoveredClip] = useState<number | null>(null);

  const [lines, setLines] = useState<{
    tl: LinePath | null;
    bl: LinePath | null;
    tr: LinePath | null;
    br: LinePath | null;
  }>({
    tl: null,
    bl: null,
    tr: null,
    br: null,
  });

  const updateLines = useCallback(() => {
    if (
      !containerRef.current ||
      !centerRef.current ||
      !clip1Ref.current ||
      !clip2Ref.current ||
      !clip3Ref.current ||
      !clip4Ref.current
    ) {
      return;
    }

    const cRect = containerRef.current.getBoundingClientRect();
    const centerRect = centerRef.current.getBoundingClientRect();
    const c1Rect = clip1Ref.current.getBoundingClientRect();
    const c2Rect = clip2Ref.current.getBoundingClientRect();
    const c3Rect = clip3Ref.current.getBoundingClientRect();
    const c4Rect = clip4Ref.current.getBoundingClientRect();

    const centerLeft = centerRect.left - cRect.left;
    const centerRight = centerRect.right - cRect.left;
    const centerUpperY = centerRect.top - cRect.top + centerRect.height * 0.3;
    const centerLowerY = centerRect.top - cRect.top + centerRect.height * 0.7;

    const c1Right = c1Rect.right - cRect.left;
    const c1Y = c1Rect.top - cRect.top + c1Rect.height * 0.5;

    const c2Right = c2Rect.right - cRect.left;
    const c2Y = c2Rect.top - cRect.top + c2Rect.height * 0.5;

    const c3Left = c3Rect.left - cRect.left;
    const c3Y = c3Rect.top - cRect.top + c3Rect.height * 0.5;

    const c4Left = c4Rect.left - cRect.left;
    const c4Y = c4Rect.top - cRect.top + c4Rect.height * 0.5;

    const midX_TL = c1Right + (centerLeft - c1Right) * 0.45;
    const pathTL = `M ${centerLeft} ${centerUpperY} C ${midX_TL} ${centerUpperY}, ${midX_TL} ${c1Y}, ${c1Right} ${c1Y}`;

    const midX_BL = c2Right + (centerLeft - c2Right) * 0.5;
    const pathBL = `M ${centerLeft} ${centerLowerY} C ${midX_BL} ${centerLowerY}, ${midX_BL} ${c2Y}, ${c2Right} ${c2Y}`;

    const midX_TR = centerRight + (c3Left - centerRight) * 0.55;
    const pathTR = `M ${centerRight} ${centerUpperY} C ${midX_TR} ${centerUpperY}, ${midX_TR} ${c3Y}, ${c3Left} ${c3Y}`;

    const midX_BR = centerRight + (c4Left - centerRight) * 0.5;
    const pathBR = `M ${centerRight} ${centerLowerY} C ${midX_BR} ${centerLowerY}, ${midX_BR} ${c4Y}, ${c4Left} ${c4Y}`;

    setLines({
      tl: {
        d: pathTL,
        startPoint: { x: centerLeft, y: centerUpperY },
        endPoint: { x: c1Right, y: c1Y },
      },
      bl: {
        d: pathBL,
        startPoint: { x: centerLeft, y: centerLowerY },
        endPoint: { x: c2Right, y: c2Y },
      },
      tr: {
        d: pathTR,
        startPoint: { x: centerRight, y: centerUpperY },
        endPoint: { x: c3Left, y: c3Y },
      },
      br: {
        d: pathBR,
        startPoint: { x: centerRight, y: centerLowerY },
        endPoint: { x: c4Left, y: c4Y },
      },
    });
  }, []);

  useEffect(() => {
    updateLines();
    const t1 = setTimeout(updateLines, 80);
    const t2 = setTimeout(updateLines, 300);

    window.addEventListener("resize", updateLines);
    return () => {
      clearTimeout(t1);
      clearTimeout(t2);
      window.removeEventListener("resize", updateLines);
    };
  }, [updateLines]);

  return (
    <div className="w-full relative select-none">
      <div
        ref={containerRef}
        className="relative max-w-6xl mx-auto px-4 min-h-[380px] lg:min-h-[420px] hidden lg:flex items-center justify-between"
      >
        <svg
          className="absolute inset-0 w-full h-full pointer-events-none z-10"
          style={{ overflow: "visible" }}
        >
          <defs>
            <filter id="darkShadow" x="-20%" y="-20%" width="140%" height="140%">
              <feDropShadow
                dx="0"
                dy="1"
                stdDeviation="2"
                floodColor="#000000"
                floodOpacity="0.25"
              />
            </filter>
          </defs>

          {lines.tl && (
            <g className="transition-opacity duration-300">
              <path
                d={lines.tl.d}
                fill="none"
                stroke={hoveredClip === 1 ? "#000000" : "#0f172a"}
                strokeWidth={hoveredClip === 1 ? 3.5 : 2.5}
                strokeLinecap="round"
                filter="url(#darkShadow)"
              />
              <circle
                cx={lines.tl.startPoint.x}
                cy={lines.tl.startPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle
                cx={lines.tl.endPoint.x}
                cy={lines.tl.endPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle r="3.5" fill="#0f172a" stroke="#ffffff" strokeWidth="1">
                <animateMotion
                  path={lines.tl.d}
                  dur="2.5s"
                  repeatCount="indefinite"
                />
              </circle>
            </g>
          )}

          {lines.bl && (
            <g className="transition-opacity duration-300">
              <path
                d={lines.bl.d}
                fill="none"
                stroke={hoveredClip === 2 ? "#000000" : "#0f172a"}
                strokeWidth={hoveredClip === 2 ? 3.5 : 2.5}
                strokeLinecap="round"
                filter="url(#darkShadow)"
              />
              <circle
                cx={lines.bl.startPoint.x}
                cy={lines.bl.startPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle
                cx={lines.bl.endPoint.x}
                cy={lines.bl.endPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle r="3.5" fill="#0f172a" stroke="#ffffff" strokeWidth="1">
                <animateMotion
                  path={lines.bl.d}
                  dur="2.8s"
                  repeatCount="indefinite"
                />
              </circle>
            </g>
          )}

          {lines.tr && (
            <g className="transition-opacity duration-300">
              <path
                d={lines.tr.d}
                fill="none"
                stroke={hoveredClip === 3 ? "#000000" : "#0f172a"}
                strokeWidth={hoveredClip === 3 ? 3.5 : 2.5}
                strokeLinecap="round"
                filter="url(#darkShadow)"
              />
              <circle
                cx={lines.tr.startPoint.x}
                cy={lines.tr.startPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle
                cx={lines.tr.endPoint.x}
                cy={lines.tr.endPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle r="3.5" fill="#0f172a" stroke="#ffffff" strokeWidth="1">
                <animateMotion
                  path={lines.tr.d}
                  dur="2.3s"
                  repeatCount="indefinite"
                />
              </circle>
            </g>
          )}

          {lines.br && (
            <g className="transition-opacity duration-300">
              <path
                d={lines.br.d}
                fill="none"
                stroke={hoveredClip === 4 ? "#000000" : "#0f172a"}
                strokeWidth={hoveredClip === 4 ? 3.5 : 2.5}
                strokeLinecap="round"
                filter="url(#darkShadow)"
              />
              <circle
                cx={lines.br.startPoint.x}
                cy={lines.br.startPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle
                cx={lines.br.endPoint.x}
                cy={lines.br.endPoint.y}
                r="4"
                fill="#0f172a"
                stroke="#ffffff"
                strokeWidth="1.5"
              />
              <circle r="3.5" fill="#0f172a" stroke="#ffffff" strokeWidth="1">
                <animateMotion
                  path={lines.br.d}
                  dur="2.7s"
                  repeatCount="indefinite"
                />
              </circle>
            </g>
          )}
        </svg>

        <div className="flex flex-col gap-6 items-end z-20 w-44 lg:w-48 xl:w-52">
          <div
            ref={clip1Ref}
            onMouseEnter={() => setHoveredClip(1)}
            onMouseLeave={() => setHoveredClip(null)}
            className={`w-full rounded-xl bg-white/40 backdrop-blur-md p-1 shadow-md transition-all duration-300 hover:scale-105 hover:shadow-xl border ${
              hoveredClip === 1 ? "border-slate-800 shadow-slate-900/20" : "border-white/60"
            }`}
          >
            <div className="relative aspect-video rounded-lg overflow-hidden">
              <img
                src={CLIPS[0].image}
                alt="Clip preview"
                className="w-full h-full object-cover"
              />
              <div className="absolute bottom-1.5 right-1.5 px-1.5 py-0.5 rounded-md bg-black/60 backdrop-blur-xs text-white text-[11px] font-mono leading-none">
                {CLIPS[0].duration}
              </div>
            </div>
          </div>

          <div
            ref={clip2Ref}
            onMouseEnter={() => setHoveredClip(2)}
            onMouseLeave={() => setHoveredClip(null)}
            className={`w-full rounded-xl bg-white/40 backdrop-blur-md p-1 shadow-md transition-all duration-300 hover:scale-105 hover:shadow-xl border ${
              hoveredClip === 2 ? "border-slate-800 shadow-slate-900/20" : "border-white/60"
            }`}
          >
            <div className="relative aspect-video rounded-lg overflow-hidden">
              <img
                src={CLIPS[1].image}
                alt="Clip preview"
                className="w-full h-full object-cover"
              />
              <div className="absolute bottom-1.5 right-1.5 px-1.5 py-0.5 rounded-md bg-black/60 backdrop-blur-xs text-white text-[11px] font-mono leading-none">
                {CLIPS[1].duration}
              </div>
            </div>
          </div>
        </div>

        <div
          ref={centerRef}
          className="relative z-20 w-[420px] lg:w-[460px] xl:w-[500px] mx-4 transition-transform duration-300"
        >
          <div className="relative rounded-2xl bg-white/30 backdrop-blur-md p-1.5 shadow-xl border-2 border-white/70">
            <div className="relative aspect-video rounded-xl overflow-hidden bg-black/80 shadow-inner group">
              <img
                src="/BIG_PODCAST.jpg"
                alt="Main video"
                className="w-full h-full object-cover"
              />

              <div className="absolute inset-0 flex items-center justify-center pointer-events-none">
                <div className="w-12 h-12 rounded-full bg-white/30 backdrop-blur-md border border-white/70 flex items-center justify-center text-white shadow-lg group-hover:scale-110 transition-transform duration-300">
                  <FaPlay className="w-4 h-4 ml-0.5 text-white" />
                </div>
              </div>

              <div className="absolute bottom-2.5 right-2.5 px-2 py-0.5 rounded-md bg-black/65 backdrop-blur-xs text-white text-xs font-mono leading-none">
                48:15
              </div>
            </div>
          </div>
        </div>

        <div className="flex flex-col gap-6 items-start z-20 w-44 lg:w-48 xl:w-52">
          <div
            ref={clip3Ref}
            onMouseEnter={() => setHoveredClip(3)}
            onMouseLeave={() => setHoveredClip(null)}
            className={`w-full rounded-xl bg-white/40 backdrop-blur-md p-1 shadow-md transition-all duration-300 hover:scale-105 hover:shadow-xl border ${
              hoveredClip === 3 ? "border-slate-800 shadow-slate-900/20" : "border-white/60"
            }`}
          >
            <div className="relative aspect-video rounded-lg overflow-hidden">
              <img
                src={CLIPS[2].image}
                alt="Clip preview"
                className="w-full h-full object-cover"
              />
              <div className="absolute bottom-1.5 right-1.5 px-1.5 py-0.5 rounded-md bg-black/60 backdrop-blur-xs text-white text-[11px] font-mono leading-none">
                {CLIPS[2].duration}
              </div>
            </div>
          </div>

          <div
            ref={clip4Ref}
            onMouseEnter={() => setHoveredClip(4)}
            onMouseLeave={() => setHoveredClip(null)}
            className={`w-full rounded-xl bg-white/40 backdrop-blur-md p-1 shadow-md transition-all duration-300 hover:scale-105 hover:shadow-xl border ${
              hoveredClip === 4 ? "border-slate-800 shadow-slate-900/20" : "border-white/60"
            }`}
          >
            <div className="relative aspect-video rounded-lg overflow-hidden">
              <img
                src={CLIPS[3].image}
                alt="Clip preview"
                className="w-full h-full object-cover"
              />
              <div className="absolute bottom-1.5 right-1.5 px-1.5 py-0.5 rounded-md bg-black/60 backdrop-blur-xs text-white text-[11px] font-mono leading-none">
                {CLIPS[3].duration}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="lg:hidden max-w-md mx-auto px-4 flex flex-col gap-4 z-20 relative">
        <div className="w-full rounded-xl bg-white/30 backdrop-blur-md p-1 shadow-lg border border-white/70">
          <div className="relative aspect-video rounded-lg overflow-hidden bg-black/80">
            <img
              src="/BIG_PODCAST.jpg"
              alt="Main video"
              className="w-full h-full object-cover"
            />
            <div className="absolute inset-0 flex items-center justify-center">
              <div className="w-10 h-10 rounded-full bg-white/30 backdrop-blur-md border border-white/60 flex items-center justify-center text-white">
                <FaPlay className="w-3.5 h-3.5 ml-0.5" />
              </div>
            </div>
            <div className="absolute bottom-2 right-2 px-1.5 py-0.5 rounded bg-black/60 text-white text-[11px] font-mono">
              48:15
            </div>
          </div>
        </div>

        <div className="grid grid-cols-2 gap-2.5">
          {CLIPS.map((clip) => (
            <div
              key={clip.id}
              className="rounded-lg bg-white/40 backdrop-blur-md p-1 shadow-sm border border-white/60"
            >
              <div className="relative aspect-video rounded-md overflow-hidden">
                <img
                  src={clip.image}
                  alt="Clip preview"
                  className="w-full h-full object-cover"
                />
                <div className="absolute bottom-1 right-1 px-1 py-0.5 rounded bg-black/60 text-white text-[10px] font-mono">
                  {clip.duration}
                </div>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}
