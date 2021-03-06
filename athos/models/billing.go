/*
 * Copyright (C) 2018 Nethesis S.r.l.
 * http://www.nethesis.it - info@nethesis.it
 *
 * This file is part of Dartagnan project.
 *
 * Dartagnan is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License,
 * or any later version.
 *
 * Dartagnan is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Dartagnan.  If not, see COPYING.
 *
 */

package models

type Billing struct {
	ID          int     `db:"id" json:"id"`
	CreatorID   string  `db:"creator_id" json:"creator_id"`
	Name        string  `db:"name" json:"name"`
	Address     string  `db:"address" json:"address"`
	Country     string  `db:"country" json:"country"`
	City        string  `db:"city" json:"city"`
	PostalCode  string  `db:"postal_code" json:"postal_code"`
	Vat         string  `db:"vat" json:"vat"`
	Tax         int     `sql:"-" json:"tax"`
}

type BillingJSON struct {
	Name        string  `db:"name" json:"name"`
	Address     string  `db:"address" json:"address"`
	Country     string  `db:"country" json:"country"`
	City        string  `db:"city" json:"city"`
	PostalCode  string  `db:"postal_code" json:"postal_code"`
	Vat         string  `db:"vat" json:"vat"`
}

