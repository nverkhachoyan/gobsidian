import type * as d3 from 'd3';
import type { SetOptional, SetRequired } from 'type-fest';
import type { Diagram } from '../Diagram.js';
import type { BaseDiagramConfig, MermaidConfig } from '../config.type.js';
export interface DiagramMetadata {
    title?: string;
    config?: MermaidConfig;
}
export interface InjectUtils {
    _log: any;
    _setLogLevel: any;
    _getConfig: any;
    _sanitizeText: any;
    _setupGraphViewbox: any;
    _commonDb: any;
    /** @deprecated as directives will be pre-processed since https://github.com/mermaid-js/mermaid/pull/4759 */
    _parseDirective: any;
}
/**
 * Generic Diagram DB that may apply to any diagram type.
 */
export interface DiagramDB {
    getConfig?: () => BaseDiagramConfig | undefined;
    clear?: () => void;
    setDiagramTitle?: (title: string) => void;
    getDiagramTitle?: () => string;
    setAccTitle?: (title: string) => void;
    getAccTitle?: () => string;
    setAccDescription?: (description: string) => void;
    getAccDescription?: () => string;
    setDisplayMode?: (title: string) => void;
    bindFunctions?: (element: Element) => void;
}
/**
 * DiagramDB with fields that is required for all new diagrams.
 */
export type DiagramDBBase<T extends BaseDiagramConfig> = {
    getConfig: () => Required<T>;
} & SetRequired<DiagramDB, 'clear' | 'getAccTitle' | 'getDiagramTitle' | 'getAccDescription' | 'setAccDescription' | 'setAccTitle' | 'setDiagramTitle'>;
export interface DiagramStyleClassDef {
    id: string;
    styles?: string[];
    textStyles?: string[];
}
export interface DiagramRenderer {
    draw: DrawDefinition;
    getClasses?: (text: string, diagram: Pick<DiagramDefinition, 'db'>) => Map<string, DiagramStyleClassDef>;
}
export interface DiagramDefinition {
    db: DiagramDB;
    renderer: DiagramRenderer;
    parser: ParserDefinition;
    styles?: any;
    init?: (config: MermaidConfig) => void;
    injectUtils?: (_log: InjectUtils['_log'], _setLogLevel: InjectUtils['_setLogLevel'], _getConfig: InjectUtils['_getConfig'], _sanitizeText: InjectUtils['_sanitizeText'], _setupGraphViewbox: InjectUtils['_setupGraphViewbox'], _commonDb: InjectUtils['_commonDb'], 
    /** @deprecated as directives will be pre-processed since https://github.com/mermaid-js/mermaid/pull/4759 */
    _parseDirective: InjectUtils['_parseDirective']) => void;
}
export interface ExternalDiagramDefinition {
    id: string;
    detector: DiagramDetector;
    loader: DiagramLoader;
}
export type DetectorRecord = SetOptional<Omit<ExternalDiagramDefinition, 'id'>, 'loader'>;
export type DiagramDetector = (text: string, config?: MermaidConfig) => boolean;
export type DiagramLoader = () => Promise<{
    id: string;
    diagram: DiagramDefinition;
}>;
/**
 * Type for function draws diagram in the tag with id: id based on the graph definition in text.
 *
 * @param text - The text of the diagram.
 * @param id - The id of the diagram which will be used as a DOM element id.
 * @param version - MermaidJS version from package.json.
 * @param diagramObject - A standard diagram containing the DB and the text and type etc of the diagram.
 */
export type DrawDefinition = (text: string, id: string, version: string, diagramObject: Diagram) => void | Promise<void>;
export interface ParserDefinition {
    parse: (text: string) => void | Promise<void>;
    parser?: {
        yy: DiagramDB;
    };
}
export type HTML = d3.Selection<HTMLIFrameElement, unknown, Element | null, unknown>;
export type SVG = d3.Selection<SVGSVGElement, unknown, Element | null, unknown>;
export type SVGGroup = d3.Selection<SVGGElement, unknown, Element | null, unknown>;
export type DiagramStylesProvider = (options?: any) => string;
