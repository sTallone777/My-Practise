package model;

import java.util.ArrayList;
import java.util.Date;
import java.util.Iterator;
import java.util.List;

public class AdjustsalaryExample {
    protected String orderByClause;

    protected boolean distinct;

    protected List<Criteria> oredCriteria;

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public AdjustsalaryExample() {
        oredCriteria = new ArrayList<Criteria>();
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public void setOrderByClause(String orderByClause) {
        this.orderByClause = orderByClause;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public String getOrderByClause() {
        return orderByClause;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public void setDistinct(boolean distinct) {
        this.distinct = distinct;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public boolean isDistinct() {
        return distinct;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public List<Criteria> getOredCriteria() {
        return oredCriteria;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public void or(Criteria criteria) {
        oredCriteria.add(criteria);
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public Criteria or() {
        Criteria criteria = createCriteriaInternal();
        oredCriteria.add(criteria);
        return criteria;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public Criteria createCriteria() {
        Criteria criteria = createCriteriaInternal();
        if (oredCriteria.size() == 0) {
            oredCriteria.add(criteria);
        }
        return criteria;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    protected Criteria createCriteriaInternal() {
        Criteria criteria = new Criteria();
        return criteria;
    }

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    public void clear() {
        oredCriteria.clear();
        orderByClause = null;
        distinct = false;
    }

    protected abstract static class GeneratedCriteria {
        protected List<Criterion> criteria;

        protected GeneratedCriteria() {
            super();
            criteria = new ArrayList<Criterion>();
        }

        public boolean isValid() {
            return criteria.size() > 0;
        }

        public List<Criterion> getAllCriteria() {
            return criteria;
        }

        public List<Criterion> getCriteria() {
            return criteria;
        }

        protected void addCriterion(String condition) {
            if (condition == null) {
                throw new RuntimeException("Value for condition cannot be null");
            }
            criteria.add(new Criterion(condition));
        }

        protected void addCriterion(String condition, Object value, String property) {
            if (value == null) {
                throw new RuntimeException("Value for " + property + " cannot be null");
            }
            criteria.add(new Criterion(condition, value));
        }

        protected void addCriterion(String condition, Object value1, Object value2, String property) {
            if (value1 == null || value2 == null) {
                throw new RuntimeException("Between values for " + property + " cannot be null");
            }
            criteria.add(new Criterion(condition, value1, value2));
        }

        protected void addCriterionForJDBCDate(String condition, Date value, String property) {
            if (value == null) {
                throw new RuntimeException("Value for " + property + " cannot be null");
            }
            addCriterion(condition, new java.sql.Date(value.getTime()), property);
        }

        protected void addCriterionForJDBCDate(String condition, List<Date> values, String property) {
            if (values == null || values.size() == 0) {
                throw new RuntimeException("Value list for " + property + " cannot be null or empty");
            }
            List<java.sql.Date> dateList = new ArrayList<java.sql.Date>();
            Iterator<Date> iter = values.iterator();
            while (iter.hasNext()) {
                dateList.add(new java.sql.Date(iter.next().getTime()));
            }
            addCriterion(condition, dateList, property);
        }

        protected void addCriterionForJDBCDate(String condition, Date value1, Date value2, String property) {
            if (value1 == null || value2 == null) {
                throw new RuntimeException("Between values for " + property + " cannot be null");
            }
            addCriterion(condition, new java.sql.Date(value1.getTime()), new java.sql.Date(value2.getTime()), property);
        }

        public Criteria andIdIsNull() {
            addCriterion("id is null");
            return (Criteria) this;
        }

        public Criteria andIdIsNotNull() {
            addCriterion("id is not null");
            return (Criteria) this;
        }

        public Criteria andIdEqualTo(Integer value) {
            addCriterion("id =", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdNotEqualTo(Integer value) {
            addCriterion("id <>", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdGreaterThan(Integer value) {
            addCriterion("id >", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdGreaterThanOrEqualTo(Integer value) {
            addCriterion("id >=", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdLessThan(Integer value) {
            addCriterion("id <", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdLessThanOrEqualTo(Integer value) {
            addCriterion("id <=", value, "id");
            return (Criteria) this;
        }

        public Criteria andIdIn(List<Integer> values) {
            addCriterion("id in", values, "id");
            return (Criteria) this;
        }

        public Criteria andIdNotIn(List<Integer> values) {
            addCriterion("id not in", values, "id");
            return (Criteria) this;
        }

        public Criteria andIdBetween(Integer value1, Integer value2) {
            addCriterion("id between", value1, value2, "id");
            return (Criteria) this;
        }

        public Criteria andIdNotBetween(Integer value1, Integer value2) {
            addCriterion("id not between", value1, value2, "id");
            return (Criteria) this;
        }

        public Criteria andEidIsNull() {
            addCriterion("eid is null");
            return (Criteria) this;
        }

        public Criteria andEidIsNotNull() {
            addCriterion("eid is not null");
            return (Criteria) this;
        }

        public Criteria andEidEqualTo(Integer value) {
            addCriterion("eid =", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidNotEqualTo(Integer value) {
            addCriterion("eid <>", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidGreaterThan(Integer value) {
            addCriterion("eid >", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidGreaterThanOrEqualTo(Integer value) {
            addCriterion("eid >=", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidLessThan(Integer value) {
            addCriterion("eid <", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidLessThanOrEqualTo(Integer value) {
            addCriterion("eid <=", value, "eid");
            return (Criteria) this;
        }

        public Criteria andEidIn(List<Integer> values) {
            addCriterion("eid in", values, "eid");
            return (Criteria) this;
        }

        public Criteria andEidNotIn(List<Integer> values) {
            addCriterion("eid not in", values, "eid");
            return (Criteria) this;
        }

        public Criteria andEidBetween(Integer value1, Integer value2) {
            addCriterion("eid between", value1, value2, "eid");
            return (Criteria) this;
        }

        public Criteria andEidNotBetween(Integer value1, Integer value2) {
            addCriterion("eid not between", value1, value2, "eid");
            return (Criteria) this;
        }

        public Criteria andAsdateIsNull() {
            addCriterion("asDate is null");
            return (Criteria) this;
        }

        public Criteria andAsdateIsNotNull() {
            addCriterion("asDate is not null");
            return (Criteria) this;
        }

        public Criteria andAsdateEqualTo(Date value) {
            addCriterionForJDBCDate("asDate =", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateNotEqualTo(Date value) {
            addCriterionForJDBCDate("asDate <>", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateGreaterThan(Date value) {
            addCriterionForJDBCDate("asDate >", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateGreaterThanOrEqualTo(Date value) {
            addCriterionForJDBCDate("asDate >=", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateLessThan(Date value) {
            addCriterionForJDBCDate("asDate <", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateLessThanOrEqualTo(Date value) {
            addCriterionForJDBCDate("asDate <=", value, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateIn(List<Date> values) {
            addCriterionForJDBCDate("asDate in", values, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateNotIn(List<Date> values) {
            addCriterionForJDBCDate("asDate not in", values, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateBetween(Date value1, Date value2) {
            addCriterionForJDBCDate("asDate between", value1, value2, "asdate");
            return (Criteria) this;
        }

        public Criteria andAsdateNotBetween(Date value1, Date value2) {
            addCriterionForJDBCDate("asDate not between", value1, value2, "asdate");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryIsNull() {
            addCriterion("beforeSalary is null");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryIsNotNull() {
            addCriterion("beforeSalary is not null");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryEqualTo(Integer value) {
            addCriterion("beforeSalary =", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryNotEqualTo(Integer value) {
            addCriterion("beforeSalary <>", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryGreaterThan(Integer value) {
            addCriterion("beforeSalary >", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryGreaterThanOrEqualTo(Integer value) {
            addCriterion("beforeSalary >=", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryLessThan(Integer value) {
            addCriterion("beforeSalary <", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryLessThanOrEqualTo(Integer value) {
            addCriterion("beforeSalary <=", value, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryIn(List<Integer> values) {
            addCriterion("beforeSalary in", values, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryNotIn(List<Integer> values) {
            addCriterion("beforeSalary not in", values, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryBetween(Integer value1, Integer value2) {
            addCriterion("beforeSalary between", value1, value2, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andBeforesalaryNotBetween(Integer value1, Integer value2) {
            addCriterion("beforeSalary not between", value1, value2, "beforesalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryIsNull() {
            addCriterion("afterSalary is null");
            return (Criteria) this;
        }

        public Criteria andAftersalaryIsNotNull() {
            addCriterion("afterSalary is not null");
            return (Criteria) this;
        }

        public Criteria andAftersalaryEqualTo(Integer value) {
            addCriterion("afterSalary =", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryNotEqualTo(Integer value) {
            addCriterion("afterSalary <>", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryGreaterThan(Integer value) {
            addCriterion("afterSalary >", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryGreaterThanOrEqualTo(Integer value) {
            addCriterion("afterSalary >=", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryLessThan(Integer value) {
            addCriterion("afterSalary <", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryLessThanOrEqualTo(Integer value) {
            addCriterion("afterSalary <=", value, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryIn(List<Integer> values) {
            addCriterion("afterSalary in", values, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryNotIn(List<Integer> values) {
            addCriterion("afterSalary not in", values, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryBetween(Integer value1, Integer value2) {
            addCriterion("afterSalary between", value1, value2, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andAftersalaryNotBetween(Integer value1, Integer value2) {
            addCriterion("afterSalary not between", value1, value2, "aftersalary");
            return (Criteria) this;
        }

        public Criteria andReasonIsNull() {
            addCriterion("reason is null");
            return (Criteria) this;
        }

        public Criteria andReasonIsNotNull() {
            addCriterion("reason is not null");
            return (Criteria) this;
        }

        public Criteria andReasonEqualTo(String value) {
            addCriterion("reason =", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonNotEqualTo(String value) {
            addCriterion("reason <>", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonGreaterThan(String value) {
            addCriterion("reason >", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonGreaterThanOrEqualTo(String value) {
            addCriterion("reason >=", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonLessThan(String value) {
            addCriterion("reason <", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonLessThanOrEqualTo(String value) {
            addCriterion("reason <=", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonLike(String value) {
            addCriterion("reason like", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonNotLike(String value) {
            addCriterion("reason not like", value, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonIn(List<String> values) {
            addCriterion("reason in", values, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonNotIn(List<String> values) {
            addCriterion("reason not in", values, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonBetween(String value1, String value2) {
            addCriterion("reason between", value1, value2, "reason");
            return (Criteria) this;
        }

        public Criteria andReasonNotBetween(String value1, String value2) {
            addCriterion("reason not between", value1, value2, "reason");
            return (Criteria) this;
        }

        public Criteria andRemarkIsNull() {
            addCriterion("remark is null");
            return (Criteria) this;
        }

        public Criteria andRemarkIsNotNull() {
            addCriterion("remark is not null");
            return (Criteria) this;
        }

        public Criteria andRemarkEqualTo(String value) {
            addCriterion("remark =", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkNotEqualTo(String value) {
            addCriterion("remark <>", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkGreaterThan(String value) {
            addCriterion("remark >", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkGreaterThanOrEqualTo(String value) {
            addCriterion("remark >=", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkLessThan(String value) {
            addCriterion("remark <", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkLessThanOrEqualTo(String value) {
            addCriterion("remark <=", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkLike(String value) {
            addCriterion("remark like", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkNotLike(String value) {
            addCriterion("remark not like", value, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkIn(List<String> values) {
            addCriterion("remark in", values, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkNotIn(List<String> values) {
            addCriterion("remark not in", values, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkBetween(String value1, String value2) {
            addCriterion("remark between", value1, value2, "remark");
            return (Criteria) this;
        }

        public Criteria andRemarkNotBetween(String value1, String value2) {
            addCriterion("remark not between", value1, value2, "remark");
            return (Criteria) this;
        }
    }

    public static class Criteria extends GeneratedCriteria {

        protected Criteria() {
            super();
        }
    }

    public static class Criterion {
        private String condition;

        private Object value;

        private Object secondValue;

        private boolean noValue;

        private boolean singleValue;

        private boolean betweenValue;

        private boolean listValue;

        private String typeHandler;

        public String getCondition() {
            return condition;
        }

        public Object getValue() {
            return value;
        }

        public Object getSecondValue() {
            return secondValue;
        }

        public boolean isNoValue() {
            return noValue;
        }

        public boolean isSingleValue() {
            return singleValue;
        }

        public boolean isBetweenValue() {
            return betweenValue;
        }

        public boolean isListValue() {
            return listValue;
        }

        public String getTypeHandler() {
            return typeHandler;
        }

        protected Criterion(String condition) {
            super();
            this.condition = condition;
            this.typeHandler = null;
            this.noValue = true;
        }

        protected Criterion(String condition, Object value, String typeHandler) {
            super();
            this.condition = condition;
            this.value = value;
            this.typeHandler = typeHandler;
            if (value instanceof List<?>) {
                this.listValue = true;
            } else {
                this.singleValue = true;
            }
        }

        protected Criterion(String condition, Object value) {
            this(condition, value, null);
        }

        protected Criterion(String condition, Object value, Object secondValue, String typeHandler) {
            super();
            this.condition = condition;
            this.value = value;
            this.secondValue = secondValue;
            this.typeHandler = typeHandler;
            this.betweenValue = true;
        }

        protected Criterion(String condition, Object value, Object secondValue) {
            this(condition, value, secondValue, null);
        }
    }
}