//////////////////////////////////////////////////////////////////////
// select.go
//////////////////////////////////////////////////////////////////////
package accounts

import (
    "database/sql"
    "log"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
)


//////////////////////////////////////////////////////////////////////
// Select from accounts table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountsColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountsColumns, error) {
    var result []*AccountsColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNTS
    bufferQuery = append(bufferQuery, query...)
    // ID
    if len(c.ID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.ID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // FirstName
    if len(c.FirstName) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.FirstName {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_FIRST_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // LastName
    if len(c.LastName) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.LastName {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_LAST_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // MiddleName
    if len(c.MiddleName) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.MiddleName {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_MIDDLE_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Name
    if len(c.Name) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Name {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // NickName
    if len(c.NickName) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.NickName {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NICK_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Nationality
    if len(c.Nationality) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Nationality {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_NATIONALITY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Email
    if len(c.Email) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Email {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_EMAIL + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // PhoneNumber
    if len(c.PhoneNumber) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.PhoneNumber {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PHONE_NUMBER + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Password
    if len(c.Password) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Password {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PASSWORD + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Status
    if len(c.Status) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Status {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_STATUS + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // PublishableToken
    if len(c.PublishableToken) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.PublishableToken {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_PUBLISHABLE_TOKEN + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // SecretToken
    if len(c.SecretToken) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.SecretToken {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_SECRET_TOKEN + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressCountry
    if len(c.AddressCountry) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressCountry {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_COUNTRY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressCity
    if len(c.AddressCity) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressCity {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_CITY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressZipCode
    if len(c.AddressZipCode) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressZipCode {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_ZIP_CODE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Address
    if len(c.Address) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Address {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressOption
    if len(c.AddressOption) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressOption {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ADDRESS_OPTION + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // CreatedAtBefore
    if !c.CreatedAtBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_CREATED_AT + "<=" + c.CreatedAtBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNTS_CREATED_AT + "<=" + c.CreatedAtBefore.String()...)
        }
    }
    // CreatedAtAfter
    if !c.CreatedAtAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_CREATED_AT + ">=" + c.CreatedAtAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNTS_CREATED_AT + ">=" + c.CreatedAtAfter.String()...)
        }
    }
    // UpdatedAtBefore
    if !c.UpdatedAtBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_UPDATED_AT + "<=" + c.UpdatedAtBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNTS_UPDATED_AT + "<=" + c.UpdatedAtBefore.String()...)
        }
    }
    // UpdatedAtAfter
    if !c.UpdatedAtAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNTS_UPDATED_AT + ">=" + c.UpdatedAtAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNTS_UPDATED_AT + ">=" + c.UpdatedAtAfter.String()...)
        }
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colID, colFirstName, colLastName, colMiddleName, colName, colNickName, colNationality, colEmail, colPhoneNumber, colPassword, colStatus, colPublishableToken, colSecretToken, colAddressCountry, colAddressCity, colAddressZipCode, colAddress, colAddressOption sql.NullString
        var colCreatedAt, colUpdatedAt sql.NullTime
        if err := rows.Scan(&colID, &colFirstName, &colLastName, &colMiddleName, &colName, &colNickName, &colNationality, &colEmail, &colPhoneNumber, &colPassword, &colStatus, &colPublishableToken, &colSecretToken, &colAddressCountry, &colAddressCity, &colAddressZipCode, &colAddress, &colAddressOption, &colCreatedAt, &colUpdatedAt); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountsColumns{
            ID: colID.String,
            FirstName: colFirstName.String,
            LastName: colLastName.String,
            MiddleName: colMiddleName.String,
            Name: colName.String,
            NickName: colNickName.String,
            Nationality: colNationality.String,
            Email: colEmail.String,
            PhoneNumber: colPhoneNumber.String,
            Password: colPassword.String,
            Status: colStatus.String,
            PublishableToken: colPublishableToken.String,
            SecretToken: colSecretToken.String,
            AddressCountry: colAddressCountry.String,
            AddressCity: colAddressCity.String,
            AddressZipCode: colAddressZipCode.String,
            Address: colAddress.String,
            AddressOption: colAddressOption.String,
            CreatedAt: colCreatedAt.Time,
            UpdatedAt: colUpdatedAt.Time,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_meta table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountMetaColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountMetaColumns, error) {
    var result []*AccountMetaColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_META
    bufferQuery = append(bufferQuery, query...)
    // AccountID
    if len(c.AccountID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.AccountID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // MetaKey
    if len(c.MetaKey) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.MetaKey {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_META_KEY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // MetaValue
    if len(c.MetaValue) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.MetaValue {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_META_VALUE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colAccountID, colMetaKey, colMetaValue sql.NullString
        if err := rows.Scan(&colAccountID, &colMetaKey, &colMetaValue); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns:= &AccountMetaColumns{
            AccountID: colAccountID.String,
            MetaKey: colMetaKey.String,
            MetaValue: colMetaValue.String,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_images table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountImagesColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountImagesColumns, error) {
    var result []*AccountImagesColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_IMAGES
    bufferQuery = append(bufferQuery, query...)
    // ID
    if len(c.ID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.ID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // Name
    if len(c.Name) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Name {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // SizeOrMore
    if c.SizeOrMore != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_SIZE + ">=" + c.SizeOrMore...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_IMAGES_SIZE + ">=" + c.SizeOrMore...)
        }
    }
    // SizeOrLess
    if c.SizeOrLess != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_SIZE + "<=" + c.SizeOrLess...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_IMAGES_SIZE + "<=" + c.SizeOrLess...)
        }
    }
    // Path
    if len(c.Path) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Path {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_PATH + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // URI
    if len(c.URI) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.URI {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_URI + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Alt
    if len(c.Alt) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Alt {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ALT + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Type
    if len(c.Type) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Type {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_TYPE + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Link
    if len(c.Link) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Link {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_LINK + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // MimeType
    if len(c.MimeType) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.MimeType {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_MIME_TYPE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // DateBefore
    if !c.DateBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_DATE + "<=" + c.DateBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_IMAGES_DATE + "<=" + c.DateBefore.String()...)
        }
    }
    // DateAfter
    if !c.DateAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_IMAGES_DATE + ">=" + c.DateAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_IMAGES_DATE + ">=" + c.DateAfter.String()...)
        }
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colID, colName, colSize, colPath, colURI, colAlt, colType, colLink, colMimeType sql.NullString
        var colDate sql.NullTime
        if err := rows.Scan(&colID, &colName, &colSize, &colPath, &colURI, &colAlt, &colType, &colLink, &colMimeType, &colDate); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountImagesColumns{
            ID: colID.String,
            Name: colName.String,
            Size: colSize.String,
            Path: colPath.String,
            URI: colURI.String,
            Alt: colAlt.String,
            Type: colType.String,
            Link: colLink.String,
            MimeType: colMimeType.String,
            Date: colDate.Time,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountImageMapColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountImageMapColumns, error) {
    var result []*AccountImageMapColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_IMAGE_MAP
    bufferQuery = append(bufferQuery, query...)
    // AccountID
    if len(c.AccountID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.AccountID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // ImageID
    if len(c.ImageID) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.ImageID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_IMAGES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colAccountID, colImageID sql.NullString
        if err := rows.Scan(&colAccountID, &colImageID); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountImageMapColumns{
            AccountID: colAccountID.String,
            ImageID: colImageID.String,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_companies table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountCompaniesColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountCompaniesColumns, error) {
    var result []*AccountCompaniesColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANIES
    bufferQuery = append(bufferQuery, query...)
    // ID
    if len(c.ID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.ID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // Name
    if len(c.Name) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Name {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // NameEn
    if len(c.NameEn) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix...)
        }
        for i, v := range c.NameEn {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_NAME_EN + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // FoundationDate
    if !c.FoundationDate.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=" + c.FoundationDate.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANIES_FOUNDATION_DATE + "=" + c.FoundationDate.String()...)
        }
    }
    // AddressCountry
    if len(c.AddressCountry) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressCountry {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_COUNTRY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressCity
    if len(c.AddressCity) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressCity {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_CITY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressZipCode
    if len(c.AddressZipCode) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressZipCode {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_ZIP_CODE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Address
    if len(c.Address) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Address {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // AddressOption
    if len(c.AddressOption) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.AddressOption {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ADDRESS_OPTION + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Industry
    if len(c.Industry) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Industry {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_INDUSTRY + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Email
    if len(c.Email) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Email {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_EMAIL + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // PhoneNumber
    if len(c.PhoneNumber) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.PhoneNumber {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_PHONE_NUMBER + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Website
    if len(c.Website) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Website {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_WEBSITE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // CreatedAtBefore
    if !c.CreatedAtBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_CREATED_AT + "<=" + c.CreatedAtBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANIES_CREATED_AT + "<=" + c.CreatedAtBefore.String()...)
        }
    }
    // CreatedAtAfter
    if !c.CreatedAtAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_CREATED_AT + ">=" + c.CreatedAtAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANIES_CREATED_AT + ">=" + c.CreatedAtAfter.String()...)
        }
    }
    // UpdatedAtBefore
    if !c.UpdatedAtBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_UPDATED_AT + "<=" + c.UpdatedAtBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANIES_UPDATED_AT + "<=" + c.UpdatedAtBefore.String()...)
        }
    }
    // UpdatedAtAfter
    if !c.UpdatedAtAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANIES_UPDATED_AT + ">=" + c.UpdatedAtAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANIES_UPDATED_AT + ">=" + c.UpdatedAtAfter.String()...)
        }
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colID, colName, colNameEn, colAddressCountry, colAddressCity, colAddressZipCode, colAddress, colAddressOption, colIndustry, colEmail, colPhoneNumber, colWebsite sql.NullString
        var colFoundationDate, colCreatedAt, colUpdatedAt sql.NullTime
        if err := rows.Scan(&colID, &colName, &colNameEn, &colFoundationDate, &colAddressCountry, &colAddressCity, &colAddressZipCode, &colAddress, &colAddressOption, &colIndustry, &colEmail, &colPhoneNumber, &colWebsite, &colCreatedAt, &colUpdatedAt); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountCompaniesColumns{
            ID: colID.String,
            Name: colName.String,
            NameEn: colNameEn.String,
            FoundationDate: colFoundationDate.Time,
            AddressCountry: colAddressCountry.String,
            AddressCity: colAddressCity.String,
            AddressZipCode: colAddressZipCode.String,
            Address: colAddress.String,
            AddressOption: colAddressOption.String,
            Industry: colIndustry.String,
            Email: colEmail.String,
            PhoneNumber: colPhoneNumber.String,
            Website: colWebsite.String,
            CreatedAt: colCreatedAt.Time,
            UpdatedAt: colUpdatedAt.Time,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_company_map table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountCompanyMapColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountCompanyMapColumns, error) {
    var result []*AccountCompanyMapColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANY_MAP
    bufferQuery = append(bufferQuery, query...)
    // AccountID
    if len(c.AccountID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.AccountID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNTS_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // CompanyID
    if len(c.CompanyID) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.CompanyID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colAccountID, colCompanyID sql.NullString
        if err := rows.Scan(&colAccountID, &colCompanyID); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountCompanyMapColumns{
            AccountID: colAccountID.String,
            CompanyID: colCompanyID.String,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_company_images table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountCompanyImagesColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountCompanyImagesColumns, error) {
    var result []*AccountCompanyImagesColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES
    bufferQuery = append(bufferQuery, query...)
    // ID
    if len(c.ID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.ID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // Name
    if len(c.Name) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Name {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_NAME + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // SizeOrMore
    if c.SizeOrMore != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_SIZE + ">=" + c.SizeOrMore...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANY_IMAGES_SIZE + ">=" + c.SizeOrMore...)
        }
    }
    // SizeOrLess
    if c.SizeOrLess != "" {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_SIZE + "<=" + c.SizeOrLess...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANY_IMAGES_SIZE + "<=" + c.SizeOrLess...)
        }
    }
    // Path
    if len(c.Path) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Path {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_PATH + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // URI
    if len(c.URI) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.URI {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_URI + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Alt
    if len(c.Alt) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Alt {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ALT + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Type
    if len(c.Type) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Type {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_TYPE + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // Link
    if len(c.Link) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.Link {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_LINK + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // MimeType
    if len(c.MimeType) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.MimeType {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_MIME_TYPE + "='" + v + "'"...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // DateBefore
    if !c.DateBefore.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_DATE + "<=" + c.DateBefore.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANY_IMAGES_DATE + "<=" + c.DateBefore.String()...)
        }
    }
    // DateAfter
    if !c.DateAfter.IsZero() {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_DATE + ">=" + c.DateAfter.String()...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + COL_ACCOUNT_COMPANY_IMAGES_DATE + ">=" + c.DateAfter.String()...)
        }
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colID, colName, colSize, colPath, colURI, colAlt, colType, colLink, colMimeType sql.NullString
        var colDate sql.NullTime
        if err := rows.Scan(&colID, &colName, &colSize, &colPath, &colURI, &colAlt, &colType, &colLink, &colMimeType, &colDate); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountCompanyImagesColumns{
            ID: colID.String,
            Name: colName.String,
            Size: colSize.String,
            Path: colPath.String,
            URI: colURI.String,
            Alt: colAlt.String,
            Type: colType.String,
            Link: colLink.String,
            MimeType: colMimeType.String,
            Date: colDate.Time,
        }
        result = append(result, columns)
    }
    return result, nil
}


//////////////////////////////////////////////////////////////////////
// Select from account_company_image_map table.
//////////////////////////////////////////////////////////////////////
func (c *SelectAccountCompanyImageMapColumns) Select(multiSearchAnd bool, orderby string, orderDesc bool, limit int, offset int) ([]*AccountCompanyImageMapColumns, error) {
    var result []*AccountCompanyImageMapColumns
    whereFlag := true
    var multiConditionPrefix string
    if multiSearchAnd {
        multiConditionPrefix = " AND "
    } else {
        multiConditionPrefix = " OR "
    }
    bufferQuery := make([]byte, 0)
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANY_IMAGE_MAP
    bufferQuery = append(bufferQuery, query...)
    // CompanyID
    if len(c.CompanyID) > 0 {
        bufferQuery = append(bufferQuery, " WHERE ("...)
        for i, v := range c.CompanyID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANIES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
        whereFlag = false
    }
    // CompanyImageID
    if len(c.CompanyImageID) > 0 {
        if whereFlag {
            bufferQuery = append(bufferQuery, " WHERE ("...)
            whereFlag = false
        } else {
            bufferQuery = append(bufferQuery, multiConditionPrefix + "("...)
        }
        for i, v := range c.CompanyImageID {
            if i > 0 {
                bufferQuery = append(bufferQuery, " OR "...)
            }
            bufferQuery = append(bufferQuery, COL_ACCOUNT_COMPANY_IMAGES_ID + "=" + v...)
        }
        bufferQuery = append(bufferQuery, ")"...)
    }
    // ORDER BY
    if orderDesc {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " DESC"...)
    } else {
        bufferQuery = append(bufferQuery, " ORDER BY " + orderby + " ASC"...)
    }
    // LIMIT
    if limit > 0 && offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset) + "," + strconv.Itoa(limit)...)
    } else if limit > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(limit)...)
    } else if offset > 0 {
        bufferQuery = append(bufferQuery, " LIMIT " + strconv.Itoa(offset)...)
    }
    // SQL
    rows, err := db.Query(string(bufferQuery[:]))
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return result, err
    }
    defer rows.Close()
    for rows.Next() {
        var colCompanyID, colCompanyImageID sql.NullString
        if err := rows.Scan(&colCompanyID, &colCompanyImageID); err != nil {
            log.Printf("[ERROR] rows.Scan() error: %s\n", err)
            return result, err
        }
        columns := &AccountCompanyImageMapColumns{
            CompanyID: colCompanyID.String,
            CompanyImageID: colCompanyImageID.String,
        }
        result = append(result, columns)
    }
    return result, nil
}




//////////////////////////////////////////////////////////////////////
// Is the unique accounts ID.
//////////////////////////////////////////////////////////////////////
func IsUniqueAccountsID(accountID string) (bool, error) {
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNTS + " WHERE " + COL_ACCOUNTS_ID + "=" + accountID
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return false, err
    }
    defer rows.Close()
    if rows.Next() {
        return false, nil
    } else {
        return true, nil
    }
}


//////////////////////////////////////////////////////////////////////
// Is the unique account images ID.
//////////////////////////////////////////////////////////////////////
func IsUniqueAccountImagesID(accountImagesID string) (bool, error) {
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_IMAGES + " WHERE " + COL_ACCOUNT_IMAGES_ID + "=" + accountImagesID
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return false, err
    }
    defer rows.Close()
    if rows.Next() {
        return false, nil
    } else {
        return true, nil
    }
}


//////////////////////////////////////////////////////////////////////
// Is the unique account companies ID.
//////////////////////////////////////////////////////////////////////
func IsUniqueAccountCompaniesID(accountCompaniesID string) (bool, error) {
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANIES + " WHERE " + COL_ACCOUNT_COMPANIES_ID + "=" + accountCompaniesID
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return false, err
    }
    defer rows.Close()
    if rows.Next() {
        return false, nil
    } else {
        return true, nil
    }
}


//////////////////////////////////////////////////////////////////////
// Is the unique account company images ID.
//////////////////////////////////////////////////////////////////////
func IsUniqueAccountCompanyImagesID(accountCompanyImagesID string) (bool, error) {
    query := "SELECT * FROM " + TABLE_NAME_ACCOUNT_COMPANY_IMAGES + " WHERE " + COL_ACCOUNT_COMPANY_IMAGES_ID + "=" + accountCompanyImagesID
    rows, err := db.Query(query)
    if err != nil {
        log.Printf("[ERROR] db.Query() error: %s\n", err)
        return false, err
    }
    defer rows.Close()
    if rows.Next() {
        return false, nil
    } else {
        return true, nil
    }
}


//////////////////////////////////////////////////////////////////////
// Is the unique email.
//////////////////////////////////////////////////////////////////////
func IsUniqueEmail(email string) bool {
    c := SelectAccountsColumns{
        Email: []string{email},
    }
    result, err := c.Select(true, COL_ACCOUNTS_ID, true, 1, 0)
    if err != nil {
        return false
    }
    if len(result) > 0 {
        return false
    } else {
        return true
    }
}


