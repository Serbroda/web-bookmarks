import {NewSpace, Space, UpdateSpace} from "../models/space";
import {ID} from "../models/base";
import {NewPage, Page, UpdatePage} from "../models/page";
import {Link, NewLink, UpdateLink} from "../models/link";

export interface SpaceService {
    getSpaces(): Promise<Space[]>;
    getSpace(id: ID): Promise<Space | undefined>;
    createSpace(body: NewSpace): Promise<Space>;
    updateSpace(id: ID, body: UpdateSpace): Promise<Space>;
    deleteSpace(id: ID): Promise<void>;
}

export interface PageService {
    getPages(spaceId: ID): Promise<Page[]>;
    getPage(id: ID): Promise<Page | undefined>;
    createPage(spaceId: ID, body: NewPage): Promise<Page>;
    updatePage(id: ID, body: UpdatePage): Promise<Space>;
    deletePage(id: ID): Promise<void>;
}

export interface LinkService {
    getLinks(pageId: ID): Promise<Link[]>;
    getLink(id: ID): Promise<Link | undefined>;
    createLink(pageId: ID, body: NewLink): Promise<Link>;
    updateLink(id: ID, body: UpdateLink): Promise<Space>;
    deleteLink(id: ID): Promise<void>;
}
