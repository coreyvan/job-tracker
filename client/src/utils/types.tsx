
export interface ICompany {
    id: string;
    name: string;
    description: string;
    website: string;
    industries: Array<string>;
    months: number;
    location: string;
    remote_possible: boolean;
}

export interface IRole {
    id: string;
    title: string;
    company: {
        id: string;
        name: string
    };
    url: string;
    technologies: Array<string>;
    pay_lower: number;
    pay_upper: number;
    location: string;
    level: string
    remote_possible: boolean;
    posted_on: string;    
}
