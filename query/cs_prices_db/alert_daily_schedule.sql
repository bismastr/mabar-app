-- name: InsertAlertDailySchedule :exec
INSERT INTO alerts_daily_schedule (item_id, discord_id, is_active)
VALUES($1, $2, true);