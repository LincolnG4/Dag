import {
  Background,
  ReactFlow,
  useNodesState,
  useEdgesState,
  Panel,
} from '@xyflow/react';
import { useRef, useCallback } from 'react';

const nodesList = [
  {
    id: '0',
    data: { label: 'Node' },
    position: { x: 0, y: 50 },
  },
];

const Board = () => {
  const reactFlowWrapper = useRef(null);
  const [nodes, setNodes, onNodesChange] = useNodesState(nodesList);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  const onAdd = useCallback(() => {
    const newNode = {
      id: (nodes.length).toString(),
      data: { label: 'Added node' },
      position: {
        x: (Math.random() - 0.5) * 400,
        y: (Math.random() - 0.5) * 400,
      },
    };
    setNodes((nds) => nds.concat(newNode));
  }, [setNodes, nodes.length]);

  return (
    <div className="wrapper" ref={reactFlowWrapper} style={{ height: 1200, width: 2400 }}>
      <ReactFlow
        nodes={nodes}
        edges={edges}
        onNodesChange={onNodesChange}
        onEdgesChange={onEdgesChange}
        fitView
        fitViewOptions={{ padding: 2 }}
      >
        <Background />
        <Panel position="top-right">
          <button className="xy-theme__button" onClick={onAdd}>
            add node
          </button>
        </Panel>
      </ReactFlow>
    </div>
  );
};

export default Board;