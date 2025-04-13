import Joi from "joi";

const HrefObject = Joi.object({
    href: Joi.string(),
});

const PaginationLinks = Joi.object({
    first: HrefObject,
    previous: HrefObject,
    current: HrefObject,
    next: HrefObject,
    last: HrefObject,
});

const Records = Joi.object({
    first: Joi.number(),
    last: Joi.number(),
    total: Joi.number(),
    offset: Joi.number(),
    limit: Joi.number(),
});

const PaginationSchema = Joi.object({
    first: Joi.number(),
    previous: Joi.number(),
    current: Joi.number(),
    next: Joi.number(),
    last: Joi.number(),
    links: PaginationLinks,
    records: Records,
});

const Meta = Joi.object({
    pagination: PaginationSchema,
});

const PaginatedResponseSchema = Joi.object({
    data: Joi.array(),
    meta: Meta,
});

export class Paginator {
    limit: number;
    offset: number;
    page_records: number;
    total_records: number;
    first_page: number = 0;
    previous_page: number = 0;
    current_page: number = 0;
    next_page: number = 0;
    last_page: number = 0;
    base_url: string;
    current_page_first: number = 0;
    current_page_end: number = 0;

    constructor(
        limit: number,
        offset: number,
        page_records: number,
        total_records: number,
        base_url: string
    ) {
        this.limit = limit;
        this.offset = offset;
        this.page_records = page_records;
        this.total_records = total_records;
        this.base_url = base_url;

        this.current_page = Math.ceil(
            (this.page_records + this.offset) / this.limit
        );

        this.last_page = Math.ceil(this.total_records / this.limit);

        if (this.total_records > 0) {
            this.first_page = 1;
        }

        if (this.current_page > 1) {
            this.previous_page = this.current_page - 1;
        }

        if (this.last_page > 1 && this.current_page < this.last_page) {
            this.next_page = this.current_page + 1;
        }

        if (page_records > 0 && this.current_page > 0) {
            this.current_page_first = (this.current_page - 1) * this.limit + 1;
        }

        if (page_records > 0 && this.current_page > 0) {
            this.current_page_end =
                (this.current_page - 1) * this.limit + this.page_records;
        }
    }

    private get_first_page() {
        return this.first_page;
    }

    private get_previous_page() {
        return this.previous_page;
    }

    private get_current_page() {
        return this.current_page;
    }

    private get_next_page() {
        return this.next_page;
    }

    private get_last_page() {
        return this.last_page;
    }

    private get_total() {
        //Get the total number of pages

        return this.last_page;
    }

    private get_count() {
        return this.total_records;
    }

    private get_url(page: number) {
        let path_params_str = "";

        if (page > 1) {
            let offset = this.limit * (page - 1);
            let path_params = { limit: this.limit, offset: this.offset };

            for (const [key, value] of Object.entries(path_params)) {
                path_params_str = path_params_str + key + "=" + value + "&";
            }

            path_params_str = "?" + path_params_str.slice(0, -1);
        }

        return this.base_url + path_params_str;
    }

    private get_page_record_start() {
        return this.current_page_first;
    }

    private get_page_record_end() {
        return this.current_page_end;
    }

    public paginate(collection: Array<any>) {
        const links = {
            first: { href: this.get_url(this.get_first_page()) },
            previous: { href: this.get_url(this.get_previous_page()) },
            current: { href: this.get_url(this.get_current_page()) },
            next: { href: this.get_url(this.get_next_page()) },
            last: { href: this.get_url(this.get_last_page()) },
        };

        const records = {
            first: this.get_page_record_start(),
            last: this.get_page_record_end(),
            total: this.get_count(),
            offset: this.offset,
            limit: this.limit,
        };

        const pagination = {
            first: this.get_first_page(),
            previous: this.get_previous_page(),
            current: this.get_current_page(),
            next: this.get_next_page(),
            last: this.get_last_page(),
            links: links,
            records: records,
        };

        const paginated_response = {
            data: collection,
            meta: { pagination: pagination },
        };

        return paginated_response;
    }
}
