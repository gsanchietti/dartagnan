/*
 * Copyright (C) 2017 Nethesis S.r.l.
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
 * author: Edoardo Spadoni <edoardo.spadoni@nethesis.it>
 */

package models

import (
	"time"
)

type Alert struct {
	ID        int       `db:"id" json:"id"`
	AlertID   string    `db:"alert_id" json:"alert_id"`
	Priority  string    `db:"priority" json:"priority"`
	Note      string    `db:"note" json:"note"`
	Status    string    `db:"status" json:"status"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
	NameI18n  string    `sql:"-" json:"namei18n"`

	System   System `json:"system"`
	SystemID int    `db:"system_id" json:"-"`
}

type AlertHistory struct {
	ID         int       `db:"id" json:"-"`
	AlertID    string    `db:"alert_id" json:"alert_id"`
	Priority   string    `db:"priority" json:"priority"`
	Resolution string    `db:"resolution" json:"resolution"`
	StatusFrom string    `db:"status_from" json:"status_from"`
	StatusTo   string    `db:"status_to" json:"status_to"`
	StartTime  time.Time `db:"start_time" json:"start_time"`
	EndTime    time.Time `db:"end_time" json:"end_time"`

	System   System `json:"system"`
	SystemID int    `db:"system_id" json:"-"`
}

type AlertJSON struct {
	SystemID string `json:"lk"`
	AlertID  string `json:"alert_id"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

type AlertNoteJSON struct {
	SystemID string `json:"system_id"`
	Note     string `json:"note"`
}
