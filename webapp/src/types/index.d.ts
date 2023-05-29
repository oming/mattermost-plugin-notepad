export type Response = {
    channelId: string;
    channelBookmark: string;
    commonBookmark: string;
};

export type Request = {
    channelId: string;
    bookmark_content: string;
}
