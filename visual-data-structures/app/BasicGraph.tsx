import React from 'react';
import { DefaultNode, Graph } from '@visx/network';

type NetworkProps = {
    width: number,
    height: number
};

interface CustomNode {
    x: number;
    y: number;
    color?: string;
};

interface CustomLink {
    source: CustomNode;
    target: CustomNode;
    dashed?: boolean;
}

const nodes: CustomNode[] = [
    { x: 100, y: 50 },
    { x: 50, y: 150 },
    { x: 150, y: 150 },
];const links: CustomLink[] = [
    { source: nodes[0], target: nodes[1] },
    { source: nodes[1], target: nodes[2] },
    { source: nodes[2], target: nodes[0], dashed: true },
  ];

const graph = {
    nodes,
    links
}

const BasicGraph = ({ width, height }: NetworkProps) => {
    const background = '#272b4d';

    return (
        <div>
            <svg width={width} height={height}>
                <rect width={width} height={height} rx={14} fill={background} />
                <Graph<CustomLink, CustomNode>
                    graph={graph}
                    top={200}
                    left={300}
                    nodeComponent={( { node: { color }}) => 
                        color ? <DefaultNode fill={color} /> : <DefaultNode />
                    }
                />
            </svg>
        </div>
    );
};

export default BasicGraph;