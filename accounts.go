//////////////////////////////////////////////////////////////////////
// accounts.go
//
// @usage
// 
//     1. Import this package.
//
//         --------------------------------------------------
//         import myAccounts "accounts"
//         --------------------------------------------------
//
//     2. Open a database using a connector.
//
//         --------------------------------------------------
//         db, err := sql.Open("driver-name", "database=test1")
//         if err != nil {
//             // Error Handling
//         }
//         --------------------------------------------------
//
//     3. Initilize this package.
//
//         --------------------------------------------------
//         myAccounts.Init(db)
//         --------------------------------------------------
//
//
// MIT License
//
// Copyright (c) 2019 noknow
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A 
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTW//ARE.
//////////////////////////////////////////////////////////////////////
package accounts

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)


var (
    db *sql.DB
)


//////////////////////////////////////////////////////////////////////
// Initialize DB tables for accounts.
//////////////////////////////////////////////////////////////////////
func Init(mydb *sql.DB) {
    db = mydb
    query := "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNTS +
            "(" + COL_ACCOUNTS_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Account ID.'," +
            COL_ACCOUNTS_FIRST_NAME + " VARCHAR(255) COMMENT 'First name.'," +
            COL_ACCOUNTS_LAST_NAME + " VARCHAR(255) COMMENT 'Last name.'," +
            COL_ACCOUNTS_MIDDLE_NAME + " VARCHAR(255) COMMENT 'Middle name.'," +
            COL_ACCOUNTS_NAME + " VARCHAR(255) NOT NULL COMMENT 'Name.'," +
            COL_ACCOUNTS_NICK_NAME + " VARCHAR(255) COMMENT 'Nick name.'," +
            COL_ACCOUNTS_NATIONALITY + " VARCHAR(2) COMMENT 'Nationality which is country code with 2 digits.'," +
            COL_ACCOUNTS_EMAIL + " VARCHAR(255) NOT NULL COMMENT 'Email address.'," +
            COL_ACCOUNTS_PHONE_NUMBER + " VARCHAR(255) COMMENT 'phone number.'," +
            COL_ACCOUNTS_PASSWORD + " VARCHAR(255) NOT NULL COMMENT 'Password.'," +
            COL_ACCOUNTS_STATUS + " TINYINT(1) UNSIGNED NOT NULL DEFAULT 10 COMMENT '10: initial stage for personal, 11: active for personal, 20: initial stage for business, 21: active for business, 30: initial stage for manager, 31: active for manager, 0: inactive'," +
            COL_ACCOUNTS_PUBLISHABLE_TOKEN + " VARCHAR(255) COMMENT 'Publishable token.'," +
            COL_ACCOUNTS_SECRET_TOKEN + " VARCHAR(255) COMMENT 'Secret token.'," +
            COL_ACCOUNTS_ADDRESS_COUNTRY + " VARCHAR(2) COMMENT 'Country which is country code with 2 digits.'," +
            COL_ACCOUNTS_ADDRESS_CITY + " VARCHAR(255) COMMENT 'City.'," +
            COL_ACCOUNTS_ADDRESS_ZIP_CODE + " VARCHAR(20) COMMENT 'Zip code.'," +
            COL_ACCOUNTS_ADDRESS + " VARCHAR(255) COMMENT 'Address.'," +
            COL_ACCOUNTS_ADDRESS_OPTION + " VARCHAR(255) COMMENT 'Address option.'," +
            COL_ACCOUNTS_CREATED_AT + " DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP," +
            COL_ACCOUNTS_UPDATED_AT + " DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
            " PRIMARY KEY(" + COL_ACCOUNTS_ID + ")" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNTS + " table.'"
    _, err := db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_META +
            "(" + COL_ACCOUNTS_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Account ID.'," +
            COL_ACCOUNT_META_KEY + " VARCHAR(255) NOT NULL COMMENT 'Meta key.'," +
            COL_ACCOUNT_META_VALUE + " VARCHAR(255) NOT NULL COMMENT 'Meta value.'," +
            " PRIMARY KEY(" + COL_ACCOUNTS_ID + ")," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_META + "_" + TABLE_NAME_ACCOUNTS + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNTS_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNTS + " (" + COL_ACCOUNTS_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_META + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_IMAGES +
            "(" + COL_ACCOUNT_IMAGES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Image ID.'," +
            COL_ACCOUNT_IMAGES_NAME + " VARCHAR(255) NOT NULL COMMENT 'Image name.'," +
            COL_ACCOUNT_IMAGES_SIZE + " INT UNSIGNED NOT NULL COMMENT 'Image size.'," +
            COL_ACCOUNT_IMAGES_PATH + " VARCHAR(255) NOT NULL COMMENT 'Image path.'," +
            COL_ACCOUNT_IMAGES_URI + " VARCHAR(255) NOT NULL COMMENT 'Image URI.'," +
            COL_ACCOUNT_IMAGES_ALT + " VARCHAR(255) COMMENT 'Image alt.'," +
            COL_ACCOUNT_IMAGES_TYPE + " TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: normal, 1: thumbnail, 2: header'," +
            COL_ACCOUNT_IMAGES_LINK + " VARCHAR(255) COMMENT 'link.'," +
            COL_ACCOUNT_IMAGES_MIME_TYPE + " VARCHAR(255) COMMENT 'Image MIME type.'," +
            COL_ACCOUNT_IMAGES_DATE + " DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Uploaded date.'," +
            " PRIMARY KEY(" + COL_ACCOUNT_IMAGES_ID + ")" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_IMAGES + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_IMAGE_MAP +
            "(" + COL_ACCOUNTS_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Account ID.'," +
            COL_ACCOUNT_IMAGES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Image ID.'," +
            " PRIMARY KEY (" + COL_ACCOUNTS_ID + "," + COL_ACCOUNT_IMAGES_ID + ")," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_IMAGE_MAP + "_" + TABLE_NAME_ACCOUNTS + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNTS_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNTS + " (" + COL_ACCOUNTS_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_IMAGE_MAP + "_" + TABLE_NAME_ACCOUNT_IMAGES + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNT_IMAGES_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNT_IMAGES + " (" + COL_ACCOUNT_IMAGES_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_IMAGE_MAP + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_COMPANIES +
            "(" + COL_ACCOUNT_COMPANIES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Company ID.'," +
            COL_ACCOUNT_COMPANIES_NAME + " VARCHAR(255) NOT NULL COMMENT 'Company name.'," +
            COL_ACCOUNT_COMPANIES_NAME_EN + " VARCHAR(255) NOT NULL COMMENT 'Company name in English.'," +
            COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + " DATETIME NOT NULL COMMENT 'foundation date.'," +
            COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + " VARCHAR(2) COMMENT 'Country which is country code with 2 digits.'," +
            COL_ACCOUNT_COMPANIES_ADDRESS_CITY + " VARCHAR(255) COMMENT 'City.'," +
            COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + " VARCHAR(20) COMMENT 'Zip code.'," +
            COL_ACCOUNT_COMPANIES_ADDRESS + " VARCHAR(255) COMMENT 'Address.'," +
            COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + " VARCHAR(255) COMMENT 'Address option.'," +
            COL_ACCOUNT_COMPANIES_INDUSTRY + " VARCHAR(255) COMMENT 'Industry.'," +
            COL_ACCOUNT_COMPANIES_EMAIL + " VARCHAR(255) COMMENT 'Email address.'," +
            COL_ACCOUNT_COMPANIES_PHONE_NUMBER + " VARCHAR(255) COMMENT 'phone number.'," +
            COL_ACCOUNT_COMPANIES_WEBSITE + " VARCHAR(255) COMMENT 'Website.'," +
            COL_ACCOUNT_COMPANIES_CREATED_AT + " DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP," +
            COL_ACCOUNT_COMPANIES_UPDATED_AT + " DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
            " PRIMARY KEY(" + COL_ACCOUNT_COMPANIES_ID + ")" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_COMPANIES + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_COMPANY_MAP +
            "(" + COL_ACCOUNTS_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Account ID.'," +
            COL_ACCOUNT_COMPANIES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Company ID.'," +
            " PRIMARY KEY (" + COL_ACCOUNTS_ID + "," + COL_ACCOUNT_COMPANIES_ID + ")," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_COMPANY_MAP + "_" + TABLE_NAME_ACCOUNTS + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNTS_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNTS + " (" + COL_ACCOUNTS_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_COMPANY_MAP + "_" + TABLE_NAME_ACCOUNT_COMPANIES + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNT_COMPANIES_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNT_COMPANIES + " (" + COL_ACCOUNT_COMPANIES_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_COMPANY_MAP + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES +
            "(" + COL_ACCOUNT_COMPANY_IMAGES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Company image ID.'," +
            COL_ACCOUNT_COMPANY_IMAGES_NAME + " VARCHAR(255) NOT NULL COMMENT 'Company image name.'," +
            COL_ACCOUNT_COMPANY_IMAGES_SIZE + " INT UNSIGNED NOT NULL COMMENT 'Company image size.'," +
            COL_ACCOUNT_COMPANY_IMAGES_PATH + " VARCHAR(255) NOT NULL COMMENT 'Company image path.'," +
            COL_ACCOUNT_COMPANY_IMAGES_URI + " VARCHAR(255) NOT NULL COMMENT 'Company image URI.'," +
            COL_ACCOUNT_COMPANY_IMAGES_ALT + " VARCHAR(255) COMMENT 'Company image alt.'," +
            COL_ACCOUNT_COMPANY_IMAGES_TYPE + " TINYINT(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '0: normal, 1: thumbnail, 2: header'," +
            COL_ACCOUNT_COMPANY_IMAGES_LINK + " VARCHAR(255) COMMENT 'Company image link.'," +
            COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + " VARCHAR(255) COMMENT 'Company image MIME type.'," +
            COL_ACCOUNT_COMPANY_IMAGES_DATE + " DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Uploaded date.'," +
            " PRIMARY KEY(" + COL_ACCOUNT_COMPANY_IMAGES_ID + ")" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
    query = "CREATE TABLE IF NOT EXISTS " + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP +
            "(" + COL_ACCOUNT_COMPANIES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Company ID.'," +
            COL_ACCOUNT_COMPANY_IMAGES_ID + " BIGINT UNSIGNED NOT NULL COMMENT 'Company image ID.'," +
            " PRIMARY KEY (" + COL_ACCOUNT_COMPANIES_ID + "," + COL_ACCOUNT_COMPANY_IMAGES_ID + ")," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP + "_" + TABLE_NAME_ACCOUNT_COMPANIES + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNT_COMPANIES_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNT_COMPANIES + " (" + COL_ACCOUNT_COMPANIES_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE," +
            " CONSTRAINT fk_" + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP + "_" + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + "_id" +
            " FOREIGN KEY (" + COL_ACCOUNT_COMPANY_IMAGES_ID + ")" +
            " REFERENCES " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + " (" + COL_ACCOUNT_COMPANY_IMAGES_ID + ")" +
            " ON DELETE CASCADE ON UPDATE CASCADE" +
            ") ENGINE=InnoDB DEFAULT CHARSET=" + CHARSET + " COMMENT='" + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP + " table.'"
    _, err = db.Exec(query)
    if err != nil {
        db.Close()
        log.Fatalf("[FATAL] db.Exec() error: %s\n", err)
        return
    }
}
