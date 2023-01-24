import { _ } from "$env/static/private";
import { add_attribute } from "svelte/internal";
import type { CreateAt } from "./CreatedAt";
import type { UpdatedAt } from "./UpdatedAt";

export type Dinner = {
    id: number;
    createdAt: CreateAt;
    updatedAt: UpdatedAt;
}

export type DinnerCreate = {
    CreatedAt?: CreateAt;
}
