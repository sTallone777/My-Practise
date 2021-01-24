package model;

import java.util.Date;

/**
 * Table: adjustsalary
 */
public class Adjustsalary {
    /**
     * Column: id
     */
    private Integer id;

    /**
     * Column: eid
     */
    private Integer eid;

    /**
     * Column: asDate
     * Remark: 调薪日期
     */
    private Date asdate;

    /**
     * Column: beforeSalary
     * Remark: 调前薪资
     */
    private Integer beforesalary;

    /**
     * Column: afterSalary
     * Remark: 调后薪资
     */
    private Integer aftersalary;

    /**
     * Column: reason
     * Remark: 调薪原因
     */
    private String reason;

    /**
     * Column: remark
     * Remark: 备注
     */
    private String remark;

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public Integer getEid() {
        return eid;
    }

    public void setEid(Integer eid) {
        this.eid = eid;
    }

    public Date getAsdate() {
        return asdate;
    }

    public void setAsdate(Date asdate) {
        this.asdate = asdate;
    }

    public Integer getBeforesalary() {
        return beforesalary;
    }

    public void setBeforesalary(Integer beforesalary) {
        this.beforesalary = beforesalary;
    }

    public Integer getAftersalary() {
        return aftersalary;
    }

    public void setAftersalary(Integer aftersalary) {
        this.aftersalary = aftersalary;
    }

    public String getReason() {
        return reason;
    }

    public void setReason(String reason) {
        this.reason = reason == null ? null : reason.trim();
    }

    public String getRemark() {
        return remark;
    }

    public void setRemark(String remark) {
        this.remark = remark == null ? null : remark.trim();
    }
}