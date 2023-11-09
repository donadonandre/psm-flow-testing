package com.andredonadon.psmflux.domain.account;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class AccountInput {

    @JsonProperty("reference_id")
    private String referenceId;

    @JsonProperty("document_number")
    private String documentNumber;

}
