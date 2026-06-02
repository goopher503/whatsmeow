-- v15 (compatible with v8+): Persist Baileys-compatible pairing ephemeral key and edge routing info
ALTER TABLE whatsmeow_device ADD COLUMN pairing_ephemeral_key bytea CHECK ( pairing_ephemeral_key IS NULL OR length(pairing_ephemeral_key) = 32 );
ALTER TABLE whatsmeow_device ADD COLUMN routing_info bytea;
