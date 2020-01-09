//////////////////////////////////////////////////////////////////////
// update.go
//////////////////////////////////////////////////////////////////////
package accounts

import (
    "log"
)


//////////////////////////////////////////////////////////////////////
// Update accounts table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountsColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNTS + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // ID
    if values.ID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=?"...)
        args = append(args, values.ID)
        multiArgs = true
    }
    // FirstName
    if values.FirstName != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_FIRST_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_FIRST_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.FirstName)
    }
    // LastName
    if values.LastName != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_LAST_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_LAST_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.LastName)
    }
    // MiddleName
    if values.MiddleName != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_MIDDLE_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_MIDDLE_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.MiddleName)
    }
    // Name
    if values.Name != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Name)
    }
    // NickName
    if values.NickName != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_NICK_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NICK_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.NickName)
    }
    // Nationality
    if values.Nationality != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_NATIONALITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NATIONALITY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Nationality)
    }
    // Email
    if values.Email != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_EMAIL + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_EMAIL + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Email)
    }
    // PhoneNumber
    if values.PhoneNumber != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_PHONE_NUMBER + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PHONE_NUMBER + "=?"...)
            multiArgs = true
        }
        args = append(args, values.PhoneNumber)
    }
    // Password
    if values.Password != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_PASSWORD + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PASSWORD + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Password)
    }
    // Status
    if values.Status != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_STATUS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_STATUS + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Status)
    }
    // PublishableToken
    if values.PublishableToken != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_PUBLISHABLE_TOKEN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PUBLISHABLE_TOKEN + "=?"...)
            multiArgs = true
        }
        args = append(args, values.PublishableToken)
    }
    // SecretToken
    if values.SecretToken != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_SECRET_TOKEN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_SECRET_TOKEN + "=?"...)
            multiArgs = true
        }
        args = append(args, values.SecretToken)
    }
    // AddressCountry
    if values.AddressCountry != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_ADDRESS_COUNTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_COUNTRY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressCountry)
    }
    // AddressCity
    if values.AddressCity != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_ADDRESS_CITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_CITY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressCity)
    }
    // AddressZipCode
    if values.AddressZipCode != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_ADDRESS_ZIP_CODE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_ZIP_CODE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressZipCode)
    }
    // Address
    if values.Address != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_ADDRESS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Address)
    }
    // AddressOption
    if values.AddressOption != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNTS_ADDRESS_OPTION + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_OPTION + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressOption)
    }
    // WHERE - ID
    whereFlag := true
    if whereCols.ID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ID + "=?"...)
        args = append(args, whereCols.ID)
        whereFlag = false
    }
    // WHERE - FirstName
    if whereCols.FirstName != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_FIRST_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_FIRST_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.FirstName)
    }
    // WHERE - LastName
    if whereCols.LastName != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_LAST_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_LAST_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.LastName)
    }
    // WHERE - MiddleName
    if whereCols.MiddleName != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_MIDDLE_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_MIDDLE_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.MiddleName)
    }
    // WHERE - Name
    if whereCols.Name != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Name)
    }
    // WHERE - NickName
    if whereCols.NickName != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_NICK_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_NICK_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.NickName)
    }
    // WHERE - Nationality
    if whereCols.Nationality != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_NATIONALITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_NATIONALITY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Nationality)
    }
    // WHERE - Email
    if whereCols.Email != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_EMAIL + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_EMAIL + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Email)
    }
    // WHERE - PhoneNumber
    if whereCols.PhoneNumber != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_PHONE_NUMBER + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_PHONE_NUMBER + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.PhoneNumber)
    }
    // WHERE - Password
    if whereCols.Password != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_PASSWORD + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_PASSWORD + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Password)
    }
    // WHERE - Status
    if whereCols.Status != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_STATUS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_STATUS + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Status)
    }
    // WHERE - PublishableToken
    if whereCols.PublishableToken != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_PUBLISHABLE_TOKEN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_PUBLISHABLE_TOKEN + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.PublishableToken)
    }
    // WHERE - SecretToken
    if whereCols.SecretToken != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_SECRET_TOKEN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_SECRET_TOKEN + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.SecretToken)
    }
    // WHERE - AddressCountry
    if whereCols.AddressCountry != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ADDRESS_COUNTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_ADDRESS_COUNTRY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressCountry)
    }
    // WHERE - AddressCity
    if whereCols.AddressCity != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ADDRESS_CITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_ADDRESS_CITY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressCity)
    }
    // WHERE - AddressZipCode
    if whereCols.AddressZipCode != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ADDRESS_ZIP_CODE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_ADDRESS_ZIP_CODE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressZipCode)
    }
    // WHERE - Address
    if whereCols.Address != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ADDRESS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_ADDRESS + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Address)
    }
    // WHERE - AddressOption
    if whereCols.AddressOption != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ADDRESS_OPTION + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNTS_ADDRESS_OPTION + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressOption)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_meta table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountMetaColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_META + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // AccountID
    if values.AccountID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=?"...)
        args = append(args, values.AccountID)
        multiArgs = true
    }
    // MetaKey
    if values.MetaKey != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_META_KEY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_META_KEY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.MetaKey)
    }
    // MetaValue
    if values.MetaValue != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_META_VALUE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_META_VALUE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.MetaValue)
    }
    // WHERE - AccountID
    whereFlag := true
    if whereCols.AccountID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ID + "=?"...)
        args = append(args, whereCols.AccountID)
        whereFlag = false
    }
    // WHERE - MetaKey
    if whereCols.MetaKey != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_META_KEY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_META_KEY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.MetaKey)
    }
    // WHERE - MetaValue
    if whereCols.MetaValue != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_META_VALUE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_META_VALUE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.MetaValue)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_images table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountImagesColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_IMAGES + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // ID
    if values.ID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ID + "=?"...)
        args = append(args, values.ID)
        multiArgs = true
    }
    // Name
    if values.Name != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Name)
    }
    // Size
    if values.Size != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_SIZE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_SIZE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Size)
    }
    // Path
    if values.Path != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_PATH + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_PATH + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Path)
    }
    // URI
    if values.URI != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_URI + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_URI + "=?"...)
            multiArgs = true
        }
        args = append(args, values.URI)
    }
    // Alt
    if values.Alt != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_ALT + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ALT + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Alt)
    }
    // Type
    if values.Type != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_TYPE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Type)
    }
    // Link
    if values.Link != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_LINK + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_LINK + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Link)
    }
    // MimeType
    if values.MimeType != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_MIME_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_MIME_TYPE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.MimeType)
    }
    // WHERE - ID
    whereFlag := true
    if whereCols.ID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_ID + "=?"...)
        args = append(args, whereCols.ID)
        whereFlag = false
    }
    // WHERE - Name
    if whereCols.Name != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Name)
    }
    // WHERE - Size
    if whereCols.Size != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_SIZE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_SIZE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Size)
    }
    // WHERE - Path
    if whereCols.Path != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_PATH + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_PATH + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Path)
    }
    // WHERE - URI
    if whereCols.URI != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_URI + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_URI + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.URI)
    }
    // WHERE - Alt
    if whereCols.Alt != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_ALT + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_ALT + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Alt)
    }
    // WHERE - Type
    if whereCols.Type != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_TYPE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Type)
    }
    // WHERE - Link
    if whereCols.Link != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_LINK + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_LINK + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Link)
    }
    // WHERE - MimeType
    if whereCols.MimeType != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_MIME_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_MIME_TYPE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.MimeType)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountImageMapColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_IMAGE_MAP + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // AccountID
    if values.AccountID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=?"...)
        args = append(args, values.AccountID)
        multiArgs = true
    }
    // ImageID
    if values.ImageID != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_IMAGES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ID + "=?"...)
            multiArgs = true
        }
        args = append(args, values.ImageID)
    }
    // WHERE - AccountID
    whereFlag := true
    if whereCols.AccountID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ID + "=?"...)
        args = append(args, whereCols.AccountID)
        whereFlag = false
    }
    // WHERE - ImageID
    if whereCols.ImageID != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_IMAGES_ID + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.ImageID)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_companies table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountCompaniesColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_COMPANIES + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // ID
    if values.ID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=?"...)
        args = append(args, values.ID)
        multiArgs = true
    }
    // Name
    if values.Name != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Name)
    }
    // NameEn
    if values.NameEn != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_NAME_EN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_NAME_EN + "=?"...)
            multiArgs = true
        }
        args = append(args, values.NameEn)
    }
    // FoundationDate
    if !values.FoundationDate.IsZero() {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.FoundationDate)
    }
    // AddressCountry
    if values.AddressCountry != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressCountry)
    }
    // AddressCity
    if values.AddressCity != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ADDRESS_CITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_CITY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressCity)
    }
    // AddressZipCode
    if values.AddressZipCode != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressZipCode)
    }
    // Address
    if values.Address != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ADDRESS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Address)
    }
    // AddressOption
    if values.AddressOption != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + "=?"...)
            multiArgs = true
        }
        args = append(args, values.AddressOption)
    }
    // Industry
    if values.Industry != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_INDUSTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_INDUSTRY + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Industry)
    }
    // Email
    if values.Email != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_EMAIL + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_EMAIL + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Email)
    }
    // PhoneNumber
    if values.PhoneNumber != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_PHONE_NUMBER + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_PHONE_NUMBER + "=?"...)
            multiArgs = true
        }
        args = append(args, values.PhoneNumber)
    }
    // Website
    if values.Website != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_WEBSITE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_WEBSITE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Website)
    }
    // WHERE - ID
    whereFlag := true
    if whereCols.ID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ID + "=?"...)
        args = append(args, whereCols.ID)
        whereFlag = false
    }
    // WHERE - Name
    if whereCols.Name != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Name)
    }
    // WHERE - NameEn
    if whereCols.NameEn != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_NAME_EN + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_NAME_EN + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.NameEn)
    }
    // WHERE - FoundationDate
    if !whereCols.FoundationDate.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=?"...)
            whereFlag = false
        }
        args = append(args, values.FoundationDate)
    }
    // WHERE - AddressCountry
    if whereCols.AddressCountry != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressCountry)
    }
    // WHERE - AddressCity
    if whereCols.AddressCity != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ADDRESS_CITY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ADDRESS_CITY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressCity)
    }
    // WHERE - AddressZipCode
    if whereCols.AddressZipCode != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressZipCode)
    }
    // WHERE - Address
    if whereCols.Address != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ADDRESS + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ADDRESS + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Address)
    }
    // WHERE - AddressOption
    if whereCols.AddressOption != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.AddressOption)
    }
    // WHERE - Industry
    if whereCols.Industry != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_INDUSTRY + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_INDUSTRY + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Industry)
    }
    // WHERE - Email
    if whereCols.Email != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_EMAIL + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_EMAIL + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Email)
    }
    // WHERE - PhoneNumber
    if whereCols.PhoneNumber != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_PHONE_NUMBER + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_PHONE_NUMBER + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.PhoneNumber)
    }
    // WHERE - Website
    if whereCols.Website != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_WEBSITE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_WEBSITE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Website)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_company_map table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountCompanyMapColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_COMPANY_MAP + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // AccountID
    if values.AccountID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=?"...)
        args = append(args, values.AccountID)
        multiArgs = true
    }
    // CompanyID
    if values.CompanyID != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANIES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=?"...)
            multiArgs = true
        }
        args = append(args, values.CompanyID)
    }
    // WHERE - AccountID
    whereFlag := true
    if whereCols.AccountID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_ID + "=?"...)
        args = append(args, whereCols.AccountID)
        whereFlag = false
    }
    // WHERE - CompanyID
    if whereCols.CompanyID != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANIES_ID + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.CompanyID)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_company_images table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountCompanyImagesColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // ID
    if values.ID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
        args = append(args, values.ID)
        multiArgs = true
    }
    // Name
    if values.Name != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_NAME + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Name)
    }
    // Size
    if values.Size != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_SIZE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_SIZE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Size)
    }
    // Path
    if values.Path != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_PATH + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_PATH + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Path)
    }
    // URI
    if values.URI != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_URI + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_URI + "=?"...)
            multiArgs = true
        }
        args = append(args, values.URI)
    }
    // Alt
    if values.Alt != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_ALT + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ALT + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Alt)
    }
    // Type
    if values.Type != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_TYPE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Type)
    }
    // Link
    if values.Link != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_LINK + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_LINK + "=?"...)
            multiArgs = true
        }
        args = append(args, values.Link)
    }
    // MimeType
    if values.MimeType != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + "=?"...)
            multiArgs = true
        }
        args = append(args, values.MimeType)
    }
    // WHERE - ID
    whereFlag := true
    if whereCols.ID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
        args = append(args, whereCols.ID)
        whereFlag = false
    }
    // WHERE - Name
    if whereCols.Name != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_NAME + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_NAME + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Name)
    }
    // WHERE - Size
    if whereCols.Size != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_SIZE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_SIZE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Size)
    }
    // WHERE - Path
    if whereCols.Path != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_PATH + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_PATH + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Path)
    }
    // WHERE - URI
    if whereCols.URI != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_URI + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_URI + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.URI)
    }
    // WHERE - Alt
    if whereCols.Alt != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_ALT + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_ALT + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Alt)
    }
    // WHERE - Type
    if whereCols.Type != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_TYPE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Type)
    }
    // WHERE - Link
    if whereCols.Link != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_LINK + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_LINK + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.Link)
    }
    // WHERE - MimeType
    if whereCols.MimeType != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.MimeType)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


//////////////////////////////////////////////////////////////////////
// Update account_company_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *UpdateAccountCompanyImageMapColumns) Update() error {
    values := c.Values
    whereCols := c.Where
    multiArgs := false
    bufferQuery := make([]byte, 0)
    query := "UPDATE " + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP + " SET "
    bufferQuery = append(bufferQuery, query...)
    args := make([]interface{}, 0)
    // CompanyID
    if values.CompanyID != "" {
        bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=?"...)
        args = append(args, values.CompanyID)
        multiArgs = true
    }
    // CompanyImageID
    if values.CompanyImageID != "" {
        if multiArgs {
            bufferQuery = append(bufferQuery, "," + COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
            multiArgs = true
        }
        args = append(args, values.CompanyImageID)
    }
    // WHERE - CompanyID
    whereFlag := true
    if whereCols.CompanyID != "" {
        bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_ID + "=?"...)
        args = append(args, whereCols.CompanyID)
        whereFlag = false
    }
    // WHERE - CompanyImageID
    if whereCols.CompanyImageID != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
        } else {
            bufferQuery = append(bufferQuery, " AND " + COL_ACCOUNT_COMPANY_IMAGES_ID + "=?"...)
            whereFlag = false
        }
        args = append(args, whereCols.CompanyImageID)
    }
    // SQL
    stmt, err := db.Prepare(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] [travel] db.Prepare() error: %s\n", err)
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(args...)
    if err != nil {
        log.Printf("[ERROR] [travel] stmt.Exec() error: %s\n", err)
        return err
    }
    return nil
}


