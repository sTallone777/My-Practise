package model;

import java.util.Date;

/**
 * Table: appraise
 */
public class Appraise {
    /**
     * Column: id
     */
    private Integer id;

    /**
     * Column: eid
     */
    private Integer eid;

    /**
     * Column: appDate
     * Remark: 考评日期
     */
    private Date appdate;

    /**
     * Column: appResult
     * Remark: 考评结果
     */
    private String appresult;

    /**
     * Column: appContent
     * Remark: 考评内容
     */
    private String appcontent;

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

    public Date getAppdate() {
        return appdate;
    }

    public void setAppdate(Date appdate) {
        this.appdate = appdate;
    }

    public String getAppresult() {
        return appresult;
    }

    public void setAppresult(String appresult) {
        this.appresult = appresult == null ? null : appresult.trim();
    }

    public String getAppcontent() {
        return appcontent;
    }

    public void setAppcontent(String appcontent) {
        this.appcontent = appcontent == null ? null : appcontent.trim();
    }

    public String getRemark() {
        return remark;
    }

    public void setRemark(String remark) {
        this.remark = remark == null ? null : remark.trim();
    }
}