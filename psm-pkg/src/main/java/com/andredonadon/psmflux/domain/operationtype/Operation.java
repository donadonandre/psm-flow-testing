package com.andredonadon.psmflux.domain.operationtype;

import java.util.HashMap;
import java.util.Map;

public class Operation {

    public static Map<Long, String> getTypes() {
        Map<Long, String> operationsType = new HashMap<>();
        operationsType.put(1L, "COMPRA A VISTA");
        operationsType.put(2L, "COMPRA PARCELADA");
        operationsType.put(3L, "SAQUE");
        operationsType.put(4L, "PAGAMENTO");

        return operationsType;
    }

}
